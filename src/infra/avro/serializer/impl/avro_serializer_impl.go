package serializer_impl

import (
	"MVC_DI/infra/avro/schema/manager"
	"MVC_DI/infra/avro/schema/mapping"
	"MVC_DI/infra/avro/serializer"
	payload_util "MVC_DI/util/payload"

	"google.golang.org/protobuf/proto"
)

type AvroSerializerImpl struct {
	SchemaMapping mapping.ISchemaMapping
	SchemaManager manager.ISchemaManager
}

func NewAvroSerializer(schemaMapping mapping.ISchemaMapping, schemaManager manager.ISchemaManager) *AvroSerializerImpl {
	return &AvroSerializerImpl{SchemaMapping: schemaMapping, SchemaManager: schemaManager}
}
func (s *AvroSerializerImpl) SerializeProtoMessage(message proto.Message) (subject string, schemaId int64, payload []byte, err error) {
	native, err := payload_util.ProtoToNative(message)
	if err != nil {
		return "", 0, nil, err
	}

	schema := s.SchemaMapping.GetSchemaByObject(message)
	codec, id, err := s.SchemaManager.GetOrLoadCodecBySchema(schema)
	if err != nil {
		return "", 0, nil, err
	}

	binary, err := codec.BinaryFromNative(nil, native)
	if err != nil {
		return "", 0, nil, err
	}

	return schema.Subject, int64(id), binary, nil
}

// INTERFACE
var _ serializer.IAvroSerializer = (*AvroSerializerImpl)(nil)
