package impl

import (
	"MVC_DI/section/permission/mapper"

	"gorm.io/gorm"
)

type PermissionMapperImpl struct{
	DB *gorm.DB
}

// INTERFACE
var _ mapper.PermissionMapper = (*PermissionMapperImpl)(nil)