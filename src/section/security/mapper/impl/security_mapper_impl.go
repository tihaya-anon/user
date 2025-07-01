package impl

import (
	"MVC_DI/section/security/mapper"

	"gorm.io/gorm"
)

type SecurityMapperImpl struct{
	DB *gorm.DB
}

// INTERFACE
var _ mapper.SecurityMapper = (*SecurityMapperImpl)(nil)