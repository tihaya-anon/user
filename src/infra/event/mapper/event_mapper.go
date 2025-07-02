package mapper

import (
	"MVC_DI/gen/proto"
	"context"
)

//go:generate mockgen -source=mapper.go -destination=..\..\..\mock\event\mapper\mapper_mock.go -package=mapper_mock
type EventMapper interface {
	// errors: refer to handler.HandleGrpcError() and handler.ValidateEventResponse()
	SubmitEvent(ctx context.Context, envelope *proto.KafkaEnvelope) error
}
