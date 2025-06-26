package test

import (
	"MVC_DI/global/infra/schema"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Schema(t *testing.T) {
	schema_ := schema.SchemaMapping.GetSchemaByMessage("AddAuthCredentialRequest")
	assert.NotNil(t, schema_)
	_, err := schema.SchemaManager.GetCodec(schema_.Subject, "")
	assert.Nil(t, err)
}
