package avro_serializer_impl

import (
	schema_manager "MVC_DI/global/infra/avro/schema/manager"
	schema_mapping "MVC_DI/global/infra/avro/schema/mapping"
	avro_serializer "MVC_DI/global/infra/avro/serializer"
	payload_util "MVC_DI/util/payload"

	"google.golang.org/protobuf/proto"
)

type AvroSerializerImpl struct {
	SchemaMapping schema_mapping.ISchemaMapping
	SchemaManager schema_manager.ISchemaManager
}

func NewAvroSerializer(schemaMapping schema_mapping.ISchemaMapping, schemaManager schema_manager.ISchemaManager) *AvroSerializerImpl {
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
var _ avro_serializer.IAvroSerializer = (*AvroSerializerImpl)(nil)
