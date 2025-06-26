// schema_manager_test.go
package schema_test

import (
	schema_mapping_impl "MVC_DI/infra/avro/schema/mapping/impl"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_SchemaMappingLoading(t *testing.T) {
	schemaMapping := schema_mapping_impl.NewSchemaMapping()
	assert.NotNil(t, schemaMapping.GetSchemas())
}
