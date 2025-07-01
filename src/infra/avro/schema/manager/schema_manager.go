package manager

import (
	"MVC_DI/infra/avro/schema/mapping"

	"github.com/linkedin/goavro/v2"
	"google.golang.org/protobuf/proto"
)

type ISchemaManager interface {
	GetOrLoadCodecByObject(object proto.Message) (*goavro.Codec, int, error)
	GetOrLoadCodecBySchema(schema *mapping.Schema) (*goavro.Codec, int, error)
	GetOrLoadCodecBySubject(subject, avscPath string) (*goavro.Codec, int, error)
}
