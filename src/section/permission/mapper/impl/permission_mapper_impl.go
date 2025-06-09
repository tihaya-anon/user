package permission_mapper_impl

import (
	permission_mapper "MVC_DI/section/permission/mapper"

	"gorm.io/gorm"
)

type PermissionMapperImpl struct{
	DB *gorm.DB
}

// INTERFACE
var _ permission_mapper.PermissionMapper = (*PermissionMapperImpl)(nil)