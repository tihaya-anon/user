package builder

import (
	"MVC_DI/section/permission/mapper"
	"MVC_DI/section/permission/mapper/impl"
	
	"gorm.io/gorm"
)

func (builder *PermissionMapperBuilder) Build() mapper.PermissionMapper {
	return builder.permissionMapperImpl
}

func (builder *PermissionMapperBuilder) WithDB(DB *gorm.DB) *PermissionMapperBuilder {
  builder.permissionMapperImpl.DB = DB
  return builder
}

// BUILDER
type PermissionMapperBuilder struct {
  isStrict bool
	permissionMapperImpl *impl.PermissionMapperImpl
}

func NewPermissionMapperBuilder() *PermissionMapperBuilder {
	return &PermissionMapperBuilder{
		permissionMapperImpl: &impl.PermissionMapperImpl{},
	}
}

func (builder *PermissionMapperBuilder) UseStrict() *PermissionMapperBuilder { 
  builder.isStrict = true
  return builder
}