package schema_manager_test

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/riferrei/srclient"
	"github.com/stretchr/testify/require"

	schema_manager_impl "MVC_DI/infra/avro/schema/manager/impl"
	schema_mapping_mock "MVC_DI/mock/infra/avro/schema/mapping"
)

func Test_GetOrLoadCodecBySubject(t *testing.T) {
	mockClient := srclient.CreateMockSchemaRegistryClient("mock://test")
	mockMapping := schema_mapping_mock.NewMockISchemaMapping(nil)

	sampleSchema := `{
		"type": "record",
		"name": "TestRecord",
		"fields": [
			{"name": "id", "type": "string"}
		]
	}`
	subject := "test-subject"
	avscFilename := "test_schema.avsc"
	avscPath := filepath.Join(os.TempDir(), avscFilename)

	err := os.WriteFile(avscPath, []byte(sampleSchema), 0644)
	require.NoError(t, err)
	defer os.Remove(avscPath)

	sm := schema_manager_impl.NewSchemaManager(mockClient, mockMapping)
	sm.InjectResourceRoot(os.TempDir())

	codec, id, err := sm.GetOrLoadCodecBySubject(subject, avscFilename)
	require.NoError(t, err)
	require.Equal(t, 1, id)
	require.NotNil(t, codec)
}
