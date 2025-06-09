package user_service_impl

import (
	user_service "MVC_DI/section/user/service"
	user_mapper "MVC_DI/section/user/mapper"
)

type UserServiceImpl struct{
	UserMapper user_mapper.UserMapper
}

// INTERFACE
var _ user_service.UserService = (*UserServiceImpl)(nil)
