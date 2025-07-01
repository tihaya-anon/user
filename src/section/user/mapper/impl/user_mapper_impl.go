package impl

import (
	"MVC_DI/section/user/mapper"

	"gorm.io/gorm"
)

type UserMapperImpl struct{
	DB *gorm.DB
}

// INTERFACE
var _ mapper.UserMapper = (*UserMapperImpl)(nil)