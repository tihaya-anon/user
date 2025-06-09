package credential_service_impl

import (
	credential_service "MVC_DI/section/credential/service"
	credential_mapper "MVC_DI/section/credential/mapper"
)

type CredentialServiceImpl struct{
	CredentialMapper credential_mapper.CredentialMapper
}

// INTERFACE
var _ credential_service.CredentialService = (*CredentialServiceImpl)(nil)
