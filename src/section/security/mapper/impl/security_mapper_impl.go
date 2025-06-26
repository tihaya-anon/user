package security_mapper_impl

import (
	security_mapper "MVC_DI/section/security/mapper"

	"gorm.io/gorm"
)

type SecurityMapperImpl struct {
	DB *gorm.DB
}

// INTERFACE
var _ security_mapper.SecurityMapper = (*SecurityMapperImpl)(nil)
