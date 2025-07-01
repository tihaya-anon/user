package envelope

import (
	"MVC_DI/gen/proto"

	google_proto "google.golang.org/protobuf/proto"
)

type EventSubmissionDto struct {
	Message      google_proto.Message
	Priority     proto.Priority
	DeliveryMode proto.DeliveryMode
	TriggerMode  proto.TriggerMode
}
