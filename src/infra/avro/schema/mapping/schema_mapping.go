package schema_mapping

import (
	"google.golang.org/protobuf/proto"
)

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
