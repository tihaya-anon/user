package impl

import (
	"MVC_DI/section/user/service"
	"MVC_DI/section/user/mapper"
)

type UserServiceImpl struct{
	UserMapper mapper.UserMapper
}

// INTERFACE
var _ service.UserService = (*UserServiceImpl)(nil)
