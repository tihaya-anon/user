package impl

import (
	"MVC_DI/config"
	"MVC_DI/infra/avro/schema/mapping"
	"reflect"

	"google.golang.org/protobuf/proto"
)

type SchemaMappingImpl struct {
	Schemas []*mapping.Schema
}

// NewSchemaMapping load from system
func NewSchemaMapping() mapping.ISchemaMapping {
	var schemaMappingImpl = &SchemaMappingImpl{}
	path := "avro/" + config.Application.Env + "/schema_registry_mapping"
	config.Parse(path, schemaMappingImpl)
	return schemaMappingImpl
}

// GetSchemas implements mapping.ISchemaMapping.
func (sm *SchemaMappingImpl) GetSchemas() []*mapping.Schema {
	return sm.Schemas
}

func (sm *SchemaMappingImpl) GetSchemaByMessage(message string) *mapping.Schema {
	for _, schema := range sm.Schemas {
		if schema.Message == message {
			return schema
		}
	}
	return nil
}

func (sm *SchemaMappingImpl) GetSchemaByObject(object proto.Message) *mapping.Schema {
	return sm.GetSchemaByMessage(getName(object))
}

func getName[T any](v T) string {
	t := reflect.TypeOf(v)

	for t.Kind() == reflect.Pointer {
		t = t.Elem()
	}

	if name := t.Name(); name != "" {
		return name
	}

	return t.String()
}

// INTERFACE
var _ mapping.ISchemaMapping = (*SchemaMappingImpl)(nil)
