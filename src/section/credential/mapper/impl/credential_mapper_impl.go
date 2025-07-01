package impl

import (
	"MVC_DI/section/credential/mapper"

	"gorm.io/gorm"
)

type CredentialMapperImpl struct{
	DB *gorm.DB
}

// INTERFACE
var _ mapper.CredentialMapper = (*CredentialMapperImpl)(nil)