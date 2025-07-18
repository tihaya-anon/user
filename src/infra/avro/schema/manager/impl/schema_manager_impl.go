package impl

import (
	"fmt"
	"os"
	"path"
	"sync"

	"github.com/linkedin/goavro/v2"
	"github.com/riferrei/srclient"
	"google.golang.org/protobuf/proto"

	"MVC_DI/global/module"
	"MVC_DI/infra/avro/schema/manager"
	"MVC_DI/infra/avro/schema/mapping"
)

// SchemaManagerImpl Manage schema loading, caching, parsing
type SchemaManagerImpl struct {
	client        srclient.ISchemaRegistryClient
	mapping       mapping.ISchemaMapping
	codecCache    map[string]*goavro.Codec
	schemaIdCache map[string]int
	mu            sync.RWMutex
	resourcePath  string
}

// NewSchemaManager constructor
func NewSchemaManager(client srclient.ISchemaRegistryClient, mapping mapping.ISchemaMapping) *SchemaManagerImpl {
	return &SchemaManagerImpl{
		client:        client,
		mapping:       mapping,
		codecCache:    make(map[string]*goavro.Codec),
		schemaIdCache: make(map[string]int),
		resourcePath:  module.GetResource(),
	}
}
func (sm *SchemaManagerImpl) InjectResourceRoot(resourceRoot string) {
	sm.resourcePath = resourceRoot
}
func (sm *SchemaManagerImpl) GetOrLoadCodecByObject(object proto.Message) (*goavro.Codec, int, error) {
	return sm.GetOrLoadCodecBySchema(sm.mapping.GetSchemaByObject(object))
}
func (sm *SchemaManagerImpl) GetOrLoadCodecBySchema(schema *mapping.Schema) (*goavro.Codec, int, error) {
	return sm.GetOrLoadCodecBySubject(schema.Subject, schema.AvscPath)
}
func (sm *SchemaManagerImpl) GetOrLoadCodecBySubject(subject, avscPath string) (*goavro.Codec, int, error) {
	sm.mu.RLock()
	if codec, ok := sm.codecCache[subject]; ok {
		schemaID := sm.schemaIdCache[subject]
		sm.mu.RUnlock()
		return codec, schemaID, nil
	}
	sm.mu.RUnlock()

	sm.mu.Lock()
	defer sm.mu.Unlock()

	if codec, ok := sm.codecCache[subject]; ok {
		schemaID := sm.schemaIdCache[subject]
		return codec, schemaID, nil
	}

	if avscPath == "" {
		for _, s := range sm.mapping.GetSchemas() {
			if s.Subject == subject {
				avscPath = s.AvscPath
				break
			}
		}
	}

	if avscPath == "" {
		return nil, 0, fmt.Errorf("schema not found for subject: %s", subject)
	}

	schemaStr, err := loadSchemaFile(path.Join(sm.resourcePath, avscPath))
	if err != nil {
		return nil, 0, err
	}
	schema, err := sm.client.CreateSchema(subject, schemaStr, srclient.Avro)
	if err != nil {
		return nil, 0, fmt.Errorf("schema register failed: %w", err)
	}
	codec, err := goavro.NewCodec(schemaStr)
	if err != nil {
		return nil, 0, fmt.Errorf("failed to create codec: %w", err)
	}

	sm.codecCache[subject] = codec
	sm.schemaIdCache[subject] = schema.ID()
	return codec, schema.ID(), nil
}
func loadSchemaFile(path string) (string, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return "", fmt.Errorf("failed to read schema file %s: %w", path, err)
	}
	return string(data), nil
}

// INTERFACE
var _ manager.ISchemaManager = (*SchemaManagerImpl)(nil)
