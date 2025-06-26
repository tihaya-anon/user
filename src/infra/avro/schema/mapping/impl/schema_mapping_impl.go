package schema_mapping_impl

import (
	"MVC_DI/config"
	schema_mapping "MVC_DI/infra/avro/schema/mapping"
	"reflect"

	"google.golang.org/protobuf/proto"
)

type SchemaMappingImpl struct {
	Schemas []*schema_mapping.Schema
}

// NewSchemaMapping load from system
func NewSchemaMapping() schema_mapping.ISchemaMapping {
	var schemaMappingImpl = &SchemaMappingImpl{}
	path := "avro/" + config.Application.Env + "/schema_registry_mapping"
	config.Parse(path, schemaMappingImpl)
	return schemaMappingImpl
}

// GetSchemas implements schema_mapping.ISchemaMapping.
func (sm *SchemaMappingImpl) GetSchemas() []*schema_mapping.Schema {
	return sm.Schemas
}

func (sm *SchemaMappingImpl) GetSchemaByMessage(message string) *schema_mapping.Schema {
	for _, schema := range sm.Schemas {
		if schema.Message == message {
			return schema
		}
	}
	return nil
}

func (sm *SchemaMappingImpl) GetSchemaByObject(object proto.Message) *schema_mapping.Schema {
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
var _ schema_mapping.ISchemaMapping = (*SchemaMappingImpl)(nil)
