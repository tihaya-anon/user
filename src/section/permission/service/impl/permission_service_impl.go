package permission_service_impl

import (
	permission_mapper "MVC_DI/section/permission/mapper"
	permission_service "MVC_DI/section/permission/service"
)

type PermissionServiceImpl struct {
	PermissionMapper permission_mapper.PermissionMapper
}

// INTERFACE
var _ permission_service.PermissionService = (*PermissionServiceImpl)(nil)
