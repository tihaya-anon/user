package security_service_impl

import (
	security_service "MVC_DI/section/security/service"
	security_mapper "MVC_DI/section/security/mapper"
)

type SecurityServiceImpl struct{
	SecurityMapper security_mapper.SecurityMapper
}

// INTERFACE
var _ security_service.SecurityService = (*SecurityServiceImpl)(nil)
