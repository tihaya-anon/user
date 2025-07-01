package builder

import (
	"MVC_DI/section/user/service"
	"MVC_DI/section/user/service/impl"
	"MVC_DI/section/user/mapper"
)

func (builder *UserServiceBuilder) Build() service.UserService {
	if builder.isStrict && builder.userServiceImpl.UserMapper == nil {
		panic("`UserMapper` is required")
	}
	return builder.userServiceImpl
}

func (builder *UserServiceBuilder) WithUserMapper(mapper mapper.UserMapper) *UserServiceBuilder {
	builder.userServiceImpl.UserMapper = mapper
	return builder
}

// BUILDER
type UserServiceBuilder struct {
  isStrict bool
	userServiceImpl *impl.UserServiceImpl
}

func NewUserServiceBuilder() *UserServiceBuilder {
	return &UserServiceBuilder{
		userServiceImpl: &impl.UserServiceImpl{},
	}
}

func (builder *UserServiceBuilder) UseStrict() *UserServiceBuilder { 
  builder.isStrict = true
  return builder
}