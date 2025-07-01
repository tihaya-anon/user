package envelope

import (
	"MVC_DI/gen/proto"
)

type EventEnvelopeFactory interface {
	Build(sub *EventSubmissionDto) (*proto.KafkaEnvelope, error)
}
