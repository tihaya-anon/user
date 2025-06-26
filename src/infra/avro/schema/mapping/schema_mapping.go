package schema_mapping

import (
	"google.golang.org/protobuf/proto"
)

//go:generate mockgen -source=schema_mapping.go -destination=../../../../mock/infra/avro/schema/mapping/schema_mapping_mock.go -package=schema_mapping_mock
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
