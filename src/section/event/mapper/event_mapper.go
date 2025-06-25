package event_mapper

import (
	"MVC_DI/gen/proto"
	"context"
)

//go:generate mockgen -source=event_mapper.go -destination=..\..\..\mock\event\mapper\event_mapper_mock.go -package=event_mapper_mock
type EventMapper interface {
	SubmitEvent(ctx context.Context, envelope *proto.KafkaEnvelope) error
}
