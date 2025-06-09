package auth_service_impl

import (
	auth_service "MVC_DI/section/auth/service"
	auth_mapper "MVC_DI/section/auth/mapper"
)

type AuthServiceImpl struct{
	AuthMapper auth_mapper.AuthMapper
}

// INTERFACE
var _ auth_service.AuthService = (*AuthServiceImpl)(nil)
