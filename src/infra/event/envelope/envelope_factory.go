package envelope

import (
	"MVC_DI/gen/proto"
)

//go:generate mockgen -source=envelope_factory.go -destination=..\..\..\mock\event\envelope\envelope_factory_mock.go -package=envelope_mock
type EventEnvelopeFactory interface {
	Build(sub *EventSubmissionDto) (*proto.KafkaEnvelope, error)
}
