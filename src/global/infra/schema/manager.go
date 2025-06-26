package schema

import (
	"fmt"
	"os"
	"path"
	"sync"

	"github.com/linkedin/goavro/v2"
	"github.com/riferrei/srclient"
	"google.golang.org/protobuf/proto"

	schema_model "MVC_DI/global/infra/schema/model"
	"MVC_DI/global/module"
)

// ISchemaManager Manage schema loading,caching,parsing
type ISchemaManager struct {
	client     *srclient.SchemaRegistryClient
	mapping    *schema_model.ISchemaMapping
	codecCache map[string]*goavro.Codec
	mu         sync.RWMutex
}

// newSchemaManager constructor
func newSchemaManager(schemaRegistryURL string, mapping *schema_model.ISchemaMapping) *ISchemaManager {
	client := srclient.CreateSchemaRegistryClient(schemaRegistryURL)

	return &ISchemaManager{
		client:     client,
		mapping:    mapping,
		codecCache: make(map[string]*goavro.Codec),
	}
}

func (sm *ISchemaManager) GetOrLoadCodecByObject(object proto.Message) (*goavro.Codec, error) {
	return sm.GetOrLoadCodecBySchema(sm.mapping.GetSchemaByObject(object))
}
func (sm *ISchemaManager) GetOrLoadCodecBySchema(schema *schema_model.Schema) (*goavro.Codec, error) {
	return sm.GetOrLoadCodecBySubject(schema.Subject, schema.AvscPath)
}
func (sm *ISchemaManager) GetOrLoadCodecBySubject(subject, avscPath string) (*goavro.Codec, error) {
	sm.mu.RLock()
	if codec, ok := sm.codecCache[subject]; ok {
		sm.mu.RUnlock()
		return codec, nil
	}
	sm.mu.RUnlock()

	sm.mu.Lock()
	defer sm.mu.Unlock()

	if codec, ok := sm.codecCache[subject]; ok {
		return codec, nil
	}

	if avscPath == "" {
		for _, s := range sm.mapping.Schemas {
			if s.Subject == subject {
				avscPath = s.AvscPath
				break
			}
		}
	}

	if avscPath == "" {
		return nil, fmt.Errorf("schema not found for subject: %s", subject)
	}

	schemaStr, err := loadSchemaFile(path.Join(module.GetResource(), avscPath))
	if err != nil {
		return nil, err
	}

	codec, err := goavro.NewCodec(schemaStr)
	if err != nil {
		return nil, fmt.Errorf("failed to create codec: %w", err)
	}

	sm.codecCache[subject] = codec
	return codec, nil
}
func loadSchemaFile(path string) (string, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return "", fmt.Errorf("failed to read schema file %s: %w", path, err)
	}
	return string(data), nil
}
