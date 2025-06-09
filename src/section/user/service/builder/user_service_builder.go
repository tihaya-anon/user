package user_service_builder

import (
	user_service "MVC_DI/section/user/service"
	user_service_impl "MVC_DI/section/user/service/impl"
	user_mapper "MVC_DI/section/user/mapper"
)

func (builder *UserServiceBuilder) Build() user_service.UserService {
	if builder.isStrict && builder.userServiceImpl.UserMapper == nil {
		panic("`UserMapper` is required")
	}
	return builder.userServiceImpl
}

func (builder *UserServiceBuilder) WithUserMapper(mapper user_mapper.UserMapper) *UserServiceBuilder {
	builder.userServiceImpl.UserMapper = mapper
	return builder
}

// BUILDER
type UserServiceBuilder struct {
  isStrict bool
	userServiceImpl *user_service_impl.UserServiceImpl
}

func NewUserServiceBuilder() *UserServiceBuilder {
	return &UserServiceBuilder{
		userServiceImpl: &user_service_impl.UserServiceImpl{},
	}
}

func (builder *UserServiceBuilder) UseStrict() *UserServiceBuilder { 
  builder.isStrict = true
  return builder
}