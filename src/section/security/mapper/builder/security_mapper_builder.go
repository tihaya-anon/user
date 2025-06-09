package security_mapper_builder

import (
	security_mapper "MVC_DI/section/security/mapper"
	security_mapper_impl "MVC_DI/section/security/mapper/impl"
	
	"gorm.io/gorm"
)

func (builder *SecurityMapperBuilder) Build() security_mapper.SecurityMapper {
	return builder.securityMapperImpl
}

func (builder *SecurityMapperBuilder) WithDB(DB *gorm.DB) *SecurityMapperBuilder {
  builder.securityMapperImpl.DB = DB
  return builder
}

// BUILDER
type SecurityMapperBuilder struct {
  isStrict bool
	securityMapperImpl *security_mapper_impl.SecurityMapperImpl
}

func NewSecurityMapperBuilder() *SecurityMapperBuilder {
	return &SecurityMapperBuilder{
		securityMapperImpl: &security_mapper_impl.SecurityMapperImpl{},
	}
}

func (builder *SecurityMapperBuilder) UseStrict() *SecurityMapperBuilder { 
  builder.isStrict = true
  return builder
}