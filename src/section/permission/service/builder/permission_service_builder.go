package builder

import (
	"MVC_DI/section/permission/service"
	"MVC_DI/section/permission/service/impl"
	"MVC_DI/section/permission/mapper"
)

func (builder *PermissionServiceBuilder) Build() service.PermissionService {
	if builder.isStrict && builder.permissionServiceImpl.PermissionMapper == nil {
		panic("`PermissionMapper` is required")
	}
	return builder.permissionServiceImpl
}

func (builder *PermissionServiceBuilder) WithPermissionMapper(mapper mapper.PermissionMapper) *PermissionServiceBuilder {
	builder.permissionServiceImpl.PermissionMapper = mapper
	return builder
}

// BUILDER
type PermissionServiceBuilder struct {
  isStrict bool
	permissionServiceImpl *impl.PermissionServiceImpl
}

func NewPermissionServiceBuilder() *PermissionServiceBuilder {
	return &PermissionServiceBuilder{
		permissionServiceImpl: &impl.PermissionServiceImpl{},
	}
}

func (builder *PermissionServiceBuilder) UseStrict() *PermissionServiceBuilder { 
  builder.isStrict = true
  return builder
}