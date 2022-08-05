package main

import (
	"fmt"
	"os"
	"time"

	cloudeventssdk "github.com/cloudevents/sdk-go/v2"
)

func main() {
	event := cloudeventssdk.NewEvent()
	event.SetID("operation-id") // some id from request body

	hostname, _ := os.Hostname()

	// /byte.builders/auditor/license_id/feature/info.ProductName/api_group/api_resource/
	// ref: https://github.com/cloudevents/spec/blob/v1.0.1/spec.md#source-1
	event.SetSource(fmt.Sprintf("/byte.builders/platform-apiserver/%s", hostname))
	// obj.getUID
	// ref: https://github.com/cloudevents/spec/blob/v1.0.1/spec.md#subject

	sub := fmt.Sprintf("/byte.builders/users/%d", 1)
	event.SetSubject(sub)
	// builders.byte.background_tasks.{created, updated, deleted}.v1
	// ref: https://github.com/cloudevents/spec/blob/v1.0.1/spec.md#type

	taskType := "builders.byte.background_tasks.install_chart.v1"
	event.SetType(taskType)
	event.SetTime(time.Now().UTC())

	// type
	// MsgId = source + id
}

// https://github.com/cloudevents/sdk-go/blob/main/protocol/nats_jetstream/v2/sender.go#L47-L63

func main__() {
	e := cloudeventssdk.NewEvent()
	e.MarshalJSON()
}
