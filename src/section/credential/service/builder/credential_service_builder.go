package builder

import (
	"MVC_DI/section/credential/service"
	"MVC_DI/section/credential/service/impl"
	"MVC_DI/section/credential/mapper"
)

func (builder *CredentialServiceBuilder) Build() service.CredentialService {
	if builder.isStrict && builder.credentialServiceImpl.CredentialMapper == nil {
		panic("`CredentialMapper` is required")
	}
	return builder.credentialServiceImpl
}

func (builder *CredentialServiceBuilder) WithCredentialMapper(mapper mapper.CredentialMapper) *CredentialServiceBuilder {
	builder.credentialServiceImpl.CredentialMapper = mapper
	return builder
}

// BUILDER
type CredentialServiceBuilder struct {
  isStrict bool
	credentialServiceImpl *impl.CredentialServiceImpl
}

func NewCredentialServiceBuilder() *CredentialServiceBuilder {
	return &CredentialServiceBuilder{
		credentialServiceImpl: &impl.CredentialServiceImpl{},
	}
}

func (builder *CredentialServiceBuilder) UseStrict() *CredentialServiceBuilder { 
  builder.isStrict = true
  return builder
}