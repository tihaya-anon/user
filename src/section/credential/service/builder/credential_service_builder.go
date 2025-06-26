package credential_service_builder

import (
	credential_mapper "MVC_DI/section/credential/mapper"
	credential_service "MVC_DI/section/credential/service"
	credential_service_impl "MVC_DI/section/credential/service/impl"
)

func (builder *CredentialServiceBuilder) Build() credential_service.CredentialService {
	if builder.isStrict && builder.credentialServiceImpl.CredentialMapper == nil {
		panic("`CredentialMapper` is required")
	}
	return builder.credentialServiceImpl
}

func (builder *CredentialServiceBuilder) WithCredentialMapper(mapper credential_mapper.CredentialMapper) *CredentialServiceBuilder {
	builder.credentialServiceImpl.CredentialMapper = mapper
	return builder
}

// BUILDER
type CredentialServiceBuilder struct {
	isStrict              bool
	credentialServiceImpl *credential_service_impl.CredentialServiceImpl
}

func NewCredentialServiceBuilder() *CredentialServiceBuilder {
	return &CredentialServiceBuilder{
		credentialServiceImpl: &credential_service_impl.CredentialServiceImpl{},
	}
}

func (builder *CredentialServiceBuilder) UseStrict() *CredentialServiceBuilder {
	builder.isStrict = true
	return builder
}
