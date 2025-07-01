// test.go
package test

import (
	"MVC_DI/infra/avro/schema/mapping/impl"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_SchemaMappingLoading(t *testing.T) {
	schemaMapping := impl.NewSchemaMapping()
	assert.NotNil(t, schemaMapping.GetSchemas())
}
