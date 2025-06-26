package security_service_impl

import (
	security_mapper "MVC_DI/section/security/mapper"
	security_service "MVC_DI/section/security/service"
)

type SecurityServiceImpl struct {
	SecurityMapper security_mapper.SecurityMapper
}

// INTERFACE
var _ security_service.SecurityService = (*SecurityServiceImpl)(nil)
