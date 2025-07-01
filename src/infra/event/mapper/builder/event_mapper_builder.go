package builder

import (
	"MVC_DI/gen/proto"
	"MVC_DI/infra/event/mapper"
	"MVC_DI/infra/event/mapper/impl"
)

func (builder *EventMapperBuilder) Build() mapper.EventMapper {
	if builder.isStrict && builder.eventMapperImpl.KafkaEventServiceClient == nil {
		panic("`KafkaEventServiceClient` is required")
	}
	return builder.eventMapperImpl
}
func (builder *EventMapperBuilder) WithKafkaEventServiceClient(client proto.KafkaEventServiceClient) *EventMapperBuilder {
	builder.eventMapperImpl.KafkaEventServiceClient = client
	return builder
}

// BUILDER
type EventMapperBuilder struct {
	isStrict        bool
	eventMapperImpl *impl.EventMapperImpl
}

func NewEventMapperBuilder() *EventMapperBuilder {
	return &EventMapperBuilder{
		eventMapperImpl: &impl.EventMapperImpl{},
	}
}

func (builder *EventMapperBuilder) UseStrict() *EventMapperBuilder {
	builder.isStrict = true
	return builder
}
