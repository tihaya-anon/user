package auth_service_builder

import (
	auth_mapper "MVC_DI/section/auth/mapper"
	auth_service "MVC_DI/section/auth/service"
	auth_service_impl "MVC_DI/section/auth/service/impl"
)

func (builder *AuthServiceBuilder) Build() auth_service.AuthService {
	if builder.isStrict && builder.authServiceImpl.AuthMapper == nil {
		panic("`AuthMapper` is required")
	}
	return builder.authServiceImpl
}

func (builder *AuthServiceBuilder) WithAuthMapper(mapper auth_mapper.AuthMapper) *AuthServiceBuilder {
	builder.authServiceImpl.AuthMapper = mapper
	return builder
}

// BUILDER

type AuthServiceBuilder struct {
	isStrict        bool
	authServiceImpl *auth_service_impl.AuthServiceImpl
}

func NewAuthServiceBuilder() *AuthServiceBuilder {
	return &AuthServiceBuilder{
		authServiceImpl: &auth_service_impl.AuthServiceImpl{},
	}
}

func (builder *AuthServiceBuilder) UseStrict() *AuthServiceBuilder {
	builder.isStrict = true
	return builder
}
