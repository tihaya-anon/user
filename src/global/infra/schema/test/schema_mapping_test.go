package test

import (
	"MVC_DI/gen/proto"
	"MVC_DI/global/infra/schema"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_SchemaMessage(t *testing.T) {
	schema_ := schema.SchemaMapping.GetSchemaByMessage("AddAuthCredentialRequest")
	assert.NotNil(t, schema_)
	_, _, err := schema.SchemaManager.GetOrLoadCodecBySchema(schema_)
	assert.Nil(t, err)
}

func Test_SchemaObject(t *testing.T) {
	request := proto.AcknowledgeEventRequest{}
	schema_ := schema.SchemaMapping.GetSchemaByObject(&request)
	assert.NotNil(t, schema_)
	_, _, err := schema.SchemaManager.GetOrLoadCodecBySchema(schema_)
	assert.Nil(t, err)
}
