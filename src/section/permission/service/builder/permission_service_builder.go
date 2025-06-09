package permission_service_builder

import (
	permission_service "MVC_DI/section/permission/service"
	permission_service_impl "MVC_DI/section/permission/service/impl"
	permission_mapper "MVC_DI/section/permission/mapper"
)

func (builder *PermissionServiceBuilder) Build() permission_service.PermissionService {
	if builder.isStrict && builder.permissionServiceImpl.PermissionMapper == nil {
		panic("`PermissionMapper` is required")
	}
	return builder.permissionServiceImpl
}

func (builder *PermissionServiceBuilder) WithPermissionMapper(mapper permission_mapper.PermissionMapper) *PermissionServiceBuilder {
	builder.permissionServiceImpl.PermissionMapper = mapper
	return builder
}

// BUILDER
type PermissionServiceBuilder struct {
  isStrict bool
	permissionServiceImpl *permission_service_impl.PermissionServiceImpl
}

func NewPermissionServiceBuilder() *PermissionServiceBuilder {
	return &PermissionServiceBuilder{
		permissionServiceImpl: &permission_service_impl.PermissionServiceImpl{},
	}
}

func (builder *PermissionServiceBuilder) UseStrict() *PermissionServiceBuilder { 
  builder.isStrict = true
  return builder
}