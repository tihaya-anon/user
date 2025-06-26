// schema_manager_test.go
package schema_test

import (
	"testing"

	"MVC_DI/gen/proto"
	"MVC_DI/global/infra/schema"

	sr_client_mock "MVC_DI/mock/global/infra/schema"

	"github.com/golang/mock/gomock"
	"github.com/riferrei/srclient"

	"github.com/stretchr/testify/assert"
)

func newTestSchemaManager(t *testing.T) *schema.ISchemaManager {
	ctrl := gomock.NewController(t)
	t.Cleanup(ctrl.Finish)

	// ① 创建 mock client
	mockCli := sr_client_mock.NewMockISchemaRegistryClient(ctrl)

	// ② 约束：任何一次 CreateSchema 调用都返回一个“空 schema”，避免网络 IO
	mockCli.
		EXPECT().
		CreateSchema(gomock.Any(), gomock.Any(), srclient.Avro).
		Return(&srclient.Schema{}, nil).
		AnyTimes()

	// ③ 注入到 SchemaManager；SchemaMapping 用全局那份即可
	return schema.NewSchemaManager(mockCli, schema.SchemaMapping)
}

func Test_GetOrLoadCodec_ByMessageName(t *testing.T) {
	mgr := newTestSchemaManager(t)

	sch := schema.SchemaMapping.GetSchemaByMessage("AddAuthCredentialRequest")
	assert.NotNil(t, sch)

	_, id, err := mgr.GetOrLoadCodecBySchema(sch)
	assert.NoError(t, err)
	assert.Equal(t, 0, id) // mock 返回的 *srclient.Schema{}，ID() 默认为 0
}

func Test_GetOrLoadCodec_ByObject(t *testing.T) {
	mgr := newTestSchemaManager(t)

	req := proto.AcknowledgeEventRequest{}
	sch := schema.SchemaMapping.GetSchemaByObject(&req)
	assert.NotNil(t, sch)

	_, id, err := mgr.GetOrLoadCodecBySchema(sch)
	assert.NoError(t, err)
	assert.Equal(t, 0, id)
}
