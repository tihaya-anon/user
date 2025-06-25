package auth_mapper_builder

import (
	"MVC_DI/gen/proto"
	auth_mapper "MVC_DI/section/auth/mapper"
	auth_mapper_impl "MVC_DI/section/auth/mapper/impl"
)

func (builder *AuthMapperBuilder) Build() auth_mapper.AuthMapper {
	if builder.isStrict && builder.authMapperImpl.AuthCredentialServiceClient == nil {
		panic("`AuthCredentialServiceClient` is required")
	}
	if builder.isStrict && builder.authMapperImpl.AuthSessionServiceClient == nil {
		panic("`AuthSessionServiceClient` is required")
	}
	return builder.authMapperImpl
}
func (builder *AuthMapperBuilder) WithAuthSessionServiceClient(client proto.AuthSessionServiceClient) *AuthMapperBuilder {
	builder.authMapperImpl.AuthSessionServiceClient = client
	return builder
}
func (builder *AuthMapperBuilder) WithAuthCredentialServiceClient(client proto.AuthCredentialServiceClient) *AuthMapperBuilder {
	builder.authMapperImpl.AuthCredentialServiceClient = client
	return builder
}

// BUILDER
type AuthMapperBuilder struct {
	isStrict       bool
	authMapperImpl *auth_mapper_impl.AuthMapperImpl
}

func NewAuthMapperBuilder() *AuthMapperBuilder {
	return &AuthMapperBuilder{
		authMapperImpl: &auth_mapper_impl.AuthMapperImpl{},
	}
}

func (builder *AuthMapperBuilder) UseStrict() *AuthMapperBuilder {
	builder.isStrict = true
	return builder
}
