package test

import (
	"MVC_DI/gen/proto"
	payload_util "MVC_DI/util/payload"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_ProtoToNative_RoundTrip(t *testing.T) {
	req := &proto.InvalidateSessionRequest{SessionId: 42}

	native, err := payload_util.ProtoToNative(req)
	require.NoError(t, err)
	require.NotEmpty(t, native)

	expected := map[string]any{"session_id": int64(42)}
	assert.Equal(t, expected, native)
}
