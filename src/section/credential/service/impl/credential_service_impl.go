package impl

import (
	"MVC_DI/section/credential/service"
	"MVC_DI/section/credential/mapper"
)

type CredentialServiceImpl struct{
	CredentialMapper mapper.CredentialMapper
}

// INTERFACE
var _ service.CredentialService = (*CredentialServiceImpl)(nil)
