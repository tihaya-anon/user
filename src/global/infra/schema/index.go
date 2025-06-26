package schema

import (
	"MVC_DI/config"
	schema_model "MVC_DI/global/infra/schema/model"
	"fmt"

	"github.com/riferrei/srclient"
)

var SchemaMapping = &schema_model.ISchemaMapping{}
var SchemaManager = &ISchemaManager{}

func init() {
	path := "avro/" + config.Application.Env + "/schema_registry_mapping"
	config.Parse(path, SchemaMapping)
	schemaRegistryURL := fmt.Sprintf("http://%s:%d", config.Application.SchemaRegistry.Host, config.Application.SchemaRegistry.Port)
	client := srclient.CreateSchemaRegistryClient(schemaRegistryURL)
	SchemaManager = NewSchemaManager(client, SchemaMapping)
}
