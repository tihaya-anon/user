package builder

import (
	"MVC_DI/section/credential/mapper"
	"MVC_DI/section/credential/mapper/impl"
	
	"gorm.io/gorm"
)

func (builder *CredentialMapperBuilder) Build() mapper.CredentialMapper {
	return builder.credentialMapperImpl
}

func (builder *CredentialMapperBuilder) WithDB(DB *gorm.DB) *CredentialMapperBuilder {
  builder.credentialMapperImpl.DB = DB
  return builder
}

// BUILDER
type CredentialMapperBuilder struct {
  isStrict bool
	credentialMapperImpl *impl.CredentialMapperImpl
}

func NewCredentialMapperBuilder() *CredentialMapperBuilder {
	return &CredentialMapperBuilder{
		credentialMapperImpl: &impl.CredentialMapperImpl{},
	}
}

func (builder *CredentialMapperBuilder) UseStrict() *CredentialMapperBuilder { 
  builder.isStrict = true
  return builder
}