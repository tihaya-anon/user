package credential_mapper_impl

import (
	credential_mapper "MVC_DI/section/credential/mapper"

	"gorm.io/gorm"
)

type CredentialMapperImpl struct {
	DB *gorm.DB
}

// INTERFACE
var _ credential_mapper.CredentialMapper = (*CredentialMapperImpl)(nil)
