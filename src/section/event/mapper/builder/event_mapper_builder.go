package event_mapper_builder

import (
	"MVC_DI/gen/proto"
	event_mapper "MVC_DI/section/event/mapper"
	event_mapper_impl "MVC_DI/section/event/mapper/impl"
)

func (builder *EventMapperBuilder) Build() event_mapper.EventMapper {
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
	eventMapperImpl *event_mapper_impl.EventMapperImpl
}

func NewEventMapperBuilder() *EventMapperBuilder {
	return &EventMapperBuilder{
		eventMapperImpl: &event_mapper_impl.EventMapperImpl{},
	}
}

func (builder *EventMapperBuilder) UseStrict() *EventMapperBuilder {
	builder.isStrict = true
	return builder
}
