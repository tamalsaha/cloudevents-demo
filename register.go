package main

import (
	"fmt"
	cloudeventssdk "github.com/cloudevents/sdk-go/v2"
	"reflect"
	"sync"
)

type TaskType string

var (
	tasks     = map[TaskType]TaskDef{}
	m         sync.RWMutex
	errorType = reflect.TypeOf((*error)(nil)).Elem()
)

type TaskDef struct {
	in reflect.Type
	f  interface{}
}

func Register(t TaskType, fn interface{}) error {
	m.Lock()
	defer m.Unlock()

	firstReturnIsErr := false

	typ := reflect.TypeOf(fn)
	if typ.Kind() != reflect.Func {
		panic(fmt.Sprintf("fn %s must be a function, found %s", typ, typ.Kind()))
	}
	switch typ.NumOut() {
	case 0:
		// nothing more to check
	case 1:
		etyp := typ.Out(0)
		if etyp.Implements(errorType) {
			firstReturnIsErr = true
		} else if reflect.New(etyp).Type().Implements(errorType) {
			panic(fmt.Sprintf("fn %s return type should be *%s to be considered an error", typ, etyp.Name()))
		}
	default:
		panic(fmt.Sprintf("fn %s has %d return values, at most 1 is allowed", typ, typ.NumOut()))
	}

	return nil
}

func MustRegister(t TaskType, fn interface{}) {
	if err := Register(t, fn); err != nil {
		panic(err)
	}
}

type Info struct {
}

func Invoke(nc nats.Conn, ev cloudeventssdk.Event) error {
	m.RLock()
	defer m.RUnlock()

	fn, ok := tasks[TaskType(ev.Type())]
	if !ok {
		return fmt.Errorf("no task registered for type %s", ev.Type())
	}
	
	ch := make(chan Info)
	go func() {
		// send via nats
	}()

	// invoke func
	fn()

	close(ch)

	return nil
}

//var map[]
