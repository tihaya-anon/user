package permission_mapper_builder

import (
	permission_mapper "MVC_DI/section/permission/mapper"
	permission_mapper_impl "MVC_DI/section/permission/mapper/impl"

	"gorm.io/gorm"
)

func (builder *PermissionMapperBuilder) Build() permission_mapper.PermissionMapper {
	return builder.permissionMapperImpl
}

func (builder *PermissionMapperBuilder) WithDB(DB *gorm.DB) *PermissionMapperBuilder {
	builder.permissionMapperImpl.DB = DB
	return builder
}

// BUILDER
type PermissionMapperBuilder struct {
	isStrict             bool
	permissionMapperImpl *permission_mapper_impl.PermissionMapperImpl
}

func NewPermissionMapperBuilder() *PermissionMapperBuilder {
	return &PermissionMapperBuilder{
		permissionMapperImpl: &permission_mapper_impl.PermissionMapperImpl{},
	}
}

func (builder *PermissionMapperBuilder) UseStrict() *PermissionMapperBuilder {
	builder.isStrict = true
	return builder
}
