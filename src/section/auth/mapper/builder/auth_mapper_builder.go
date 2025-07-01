package builder

import (
	"MVC_DI/gen/proto"
	"MVC_DI/section/auth/mapper"
	"MVC_DI/section/auth/mapper/impl"
)

func (builder *AuthMapperBuilder) Build() mapper.AuthMapper {
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
	authMapperImpl *impl.AuthMapperImpl
}

func NewAuthMapperBuilder() *AuthMapperBuilder {
	return &AuthMapperBuilder{
		authMapperImpl: &impl.AuthMapperImpl{},
	}
}

func (builder *AuthMapperBuilder) UseStrict() *AuthMapperBuilder {
	builder.isStrict = true
	return builder
}
