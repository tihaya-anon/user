package mapping

import (
	"google.golang.org/protobuf/proto"
)

//go:generate mockgen -source=mapping.go -destination=../../../../mock/infra/avro/schema/mapping/mapping_mock.go -package=mapping_mock
type ISchemaMapping interface {
	GetSchemaByMessage(message string) *Schema
	GetSchemaByObject(object proto.Message) *Schema
	GetSchemas() []*Schema
}

type Schema struct {
	Proto    string
	Message  string
	AvscPath string
	Subject  string
}
