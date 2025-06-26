package schema_manager

import (
	schema_mapping "MVC_DI/infra/avro/schema/mapping"

	"github.com/linkedin/goavro/v2"
	"google.golang.org/protobuf/proto"
)

type ISchemaManager interface {
	GetOrLoadCodecByObject(object proto.Message) (*goavro.Codec, int, error)
	GetOrLoadCodecBySchema(schema *schema_mapping.Schema) (*goavro.Codec, int, error)
	GetOrLoadCodecBySubject(subject, avscPath string) (*goavro.Codec, int, error)
}
