package security_service_builder

import (
	security_mapper "MVC_DI/section/security/mapper"
	security_service "MVC_DI/section/security/service"
	security_service_impl "MVC_DI/section/security/service/impl"
)

func (builder *SecurityServiceBuilder) Build() security_service.SecurityService {
	if builder.isStrict && builder.securityServiceImpl.SecurityMapper == nil {
		panic("`SecurityMapper` is required")
	}
	return builder.securityServiceImpl
}

func (builder *SecurityServiceBuilder) WithSecurityMapper(mapper security_mapper.SecurityMapper) *SecurityServiceBuilder {
	builder.securityServiceImpl.SecurityMapper = mapper
	return builder
}

// BUILDER
type SecurityServiceBuilder struct {
	isStrict            bool
	securityServiceImpl *security_service_impl.SecurityServiceImpl
}

func NewSecurityServiceBuilder() *SecurityServiceBuilder {
	return &SecurityServiceBuilder{
		securityServiceImpl: &security_service_impl.SecurityServiceImpl{},
	}
}

func (builder *SecurityServiceBuilder) UseStrict() *SecurityServiceBuilder {
	builder.isStrict = true
	return builder
}
