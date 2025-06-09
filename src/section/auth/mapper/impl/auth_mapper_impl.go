package auth_mapper_impl

import (
	auth_mapper "MVC_DI/section/auth/mapper"

	"gorm.io/gorm"
)

type AuthMapperImpl struct{
	DB *gorm.DB
}

// INTERFACE
var _ auth_mapper.AuthMapper = (*AuthMapperImpl)(nil)