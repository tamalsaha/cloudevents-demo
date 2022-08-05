package main

import (
	"fmt"
	"time"

	cloudeventssdk "github.com/cloudevents/sdk-go/v2"
)

func main() {
	event := cloudeventssdk.NewEvent()
	event.SetID("operation-id") // some id from request body

	// /byte.builders/auditor/license_id/feature/info.ProductName/api_group/api_resource/
	// ref: https://github.com/cloudevents/spec/blob/v1.0.1/spec.md#source-1
	event.SetSource(fmt.Sprintf("/byte.builders/auditor/%s/feature/%s/%s/%s", ev.LicenseID, info.ProductName, ev.ResourceID.Group, ev.ResourceID.Name))
	// obj.getUID
	// ref: https://github.com/cloudevents/spec/blob/v1.0.1/spec.md#subject
	event.SetSubject(string(ev.Resource.GetUID()))
	// builders.byte.auditor.{created, updated, deleted}.v1
	// ref: https://github.com/cloudevents/spec/blob/v1.0.1/spec.md#type
	event.SetType(string(et))
	event.SetTime(time.Now().UTC())

}
