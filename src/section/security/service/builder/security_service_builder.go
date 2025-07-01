package builder

import (
	"MVC_DI/section/security/service"
	"MVC_DI/section/security/service/impl"
	"MVC_DI/section/security/mapper"
)

func (builder *SecurityServiceBuilder) Build() service.SecurityService {
	if builder.isStrict && builder.securityServiceImpl.SecurityMapper == nil {
		panic("`SecurityMapper` is required")
	}
	return builder.securityServiceImpl
}

func (builder *SecurityServiceBuilder) WithSecurityMapper(mapper mapper.SecurityMapper) *SecurityServiceBuilder {
	builder.securityServiceImpl.SecurityMapper = mapper
	return builder
}

// BUILDER
type SecurityServiceBuilder struct {
  isStrict bool
	securityServiceImpl *impl.SecurityServiceImpl
}

func NewSecurityServiceBuilder() *SecurityServiceBuilder {
	return &SecurityServiceBuilder{
		securityServiceImpl: &impl.SecurityServiceImpl{},
	}
}

func (builder *SecurityServiceBuilder) UseStrict() *SecurityServiceBuilder { 
  builder.isStrict = true
  return builder
}