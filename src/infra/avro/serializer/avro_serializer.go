package avro_serializer

import (
	"google.golang.org/protobuf/proto"
)

type IAvroSerializer interface {
	SerializeProtoMessage(message proto.Message) (subject string, schemaId int64, payload []byte, err error)
}
