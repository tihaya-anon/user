package impl

import (
	"MVC_DI/gen/proto"
	"MVC_DI/infra/avro/serializer"
	"MVC_DI/infra/event/envelope"
)

type EventEnvelopeFactoryImpl struct {
	AvroSerializer serializer.IAvroSerializer
}

// Build implements envelope.EventEnvelopeFactory.
func (e *EventEnvelopeFactoryImpl) Build(sub *envelope.EventSubmissionDto) (*proto.KafkaEnvelope, error) {
	subject, schemaId, payload, err := e.AvroSerializer.SerializeProtoMessage(sub.Message)
	if err != nil {
		return nil, err
	}
	envelope := &proto.KafkaEnvelope{
		SchemaSubject:        subject,
		SchemaId:             schemaId,
		Priority:             sub.Priority,
		Payload:              payload,
		DeliveryMode:         sub.DeliveryMode,
		TriggerModeRequested: sub.TriggerMode,
	}
	return envelope, err
}

// INTERFACE
var _ envelope.EventEnvelopeFactory = (*EventEnvelopeFactoryImpl)(nil)
