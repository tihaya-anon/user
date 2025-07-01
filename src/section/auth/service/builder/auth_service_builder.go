package service_builder

import (
	"MVC_DI/section/auth/mapper"
	"MVC_DI/section/auth/service"
	"MVC_DI/section/auth/service/impl"
)

func (builder *AuthServiceBuilder) Build() service.AuthService {
	if builder.isStrict && builder.authServiceImpl.AuthMapper == nil {
		panic("`AuthMapper` is required")
	}
	return builder.authServiceImpl
}

func (builder *AuthServiceBuilder) WithAuthMapper(mapper mapper.AuthMapper) *AuthServiceBuilder {
	builder.authServiceImpl.AuthMapper = mapper
	return builder
}

// BUILDER

type AuthServiceBuilder struct {
	isStrict        bool
	authServiceImpl *impl.AuthServiceImpl
}

func NewAuthServiceBuilder() *AuthServiceBuilder {
	return &AuthServiceBuilder{
		authServiceImpl: &impl.AuthServiceImpl{},
	}
}

func (builder *AuthServiceBuilder) UseStrict() *AuthServiceBuilder {
	builder.isStrict = true
	return builder
}
