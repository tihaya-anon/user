package impl

import (
	"MVC_DI/section/permission/service"
	"MVC_DI/section/permission/mapper"
)

type PermissionServiceImpl struct{
	PermissionMapper mapper.PermissionMapper
}

// INTERFACE
var _ service.PermissionService = (*PermissionServiceImpl)(nil)
