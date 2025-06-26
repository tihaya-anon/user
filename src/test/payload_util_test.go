package test

import (
	"MVC_DI/gen/proto"
	"MVC_DI/global/infra/schema"
	payload_util "MVC_DI/util/payload"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_ProtoToAvroBytes_RoundTrip(t *testing.T) {
	req := &proto.InvalidateSessionRequest{SessionId: 42}

	payload, err := payload_util.ProtoToAvroBytes(req)
	require.NoError(t, err)
	require.NotEmpty(t, payload)

	codec, _ := schema.SchemaManager.GetOrLoadCodecByObject(req)
	native, _, err := codec.NativeFromBinary(payload)
	require.NoError(t, err)

	expected := map[string]any{"session_id": int64(42)}
	assert.Equal(t, expected, native)
}
