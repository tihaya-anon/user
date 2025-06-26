package user_service_impl

import (
	user_mapper "MVC_DI/section/user/mapper"
	user_service "MVC_DI/section/user/service"
)

type UserServiceImpl struct {
	UserMapper user_mapper.UserMapper
}

// INTERFACE
var _ user_service.UserService = (*UserServiceImpl)(nil)
