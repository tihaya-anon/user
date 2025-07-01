package impl

import (
	"MVC_DI/section/security/service"
	"MVC_DI/section/security/mapper"
)

type SecurityServiceImpl struct{
	SecurityMapper mapper.SecurityMapper
}

// INTERFACE
var _ service.SecurityService = (*SecurityServiceImpl)(nil)
