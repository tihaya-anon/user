package credential_mapper_builder

import (
	credential_mapper "MVC_DI/section/credential/mapper"
	credential_mapper_impl "MVC_DI/section/credential/mapper/impl"
	
	"gorm.io/gorm"
)

func (builder *CredentialMapperBuilder) Build() credential_mapper.CredentialMapper {
	return builder.credentialMapperImpl
}

func (builder *CredentialMapperBuilder) WithDB(DB *gorm.DB) *CredentialMapperBuilder {
  builder.credentialMapperImpl.DB = DB
  return builder
}

// BUILDER
type CredentialMapperBuilder struct {
  isStrict bool
	credentialMapperImpl *credential_mapper_impl.CredentialMapperImpl
}

func NewCredentialMapperBuilder() *CredentialMapperBuilder {
	return &CredentialMapperBuilder{
		credentialMapperImpl: &credential_mapper_impl.CredentialMapperImpl{},
	}
}

func (builder *CredentialMapperBuilder) UseStrict() *CredentialMapperBuilder { 
  builder.isStrict = true
  return builder
}