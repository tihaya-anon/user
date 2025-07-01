package builder

import (
	"MVC_DI/section/security/mapper"
	"MVC_DI/section/security/mapper/impl"
	
	"gorm.io/gorm"
)

func (builder *SecurityMapperBuilder) Build() mapper.SecurityMapper {
	return builder.securityMapperImpl
}

func (builder *SecurityMapperBuilder) WithDB(DB *gorm.DB) *SecurityMapperBuilder {
  builder.securityMapperImpl.DB = DB
  return builder
}

// BUILDER
type SecurityMapperBuilder struct {
  isStrict bool
	securityMapperImpl *impl.SecurityMapperImpl
}

func NewSecurityMapperBuilder() *SecurityMapperBuilder {
	return &SecurityMapperBuilder{
		securityMapperImpl: &impl.SecurityMapperImpl{},
	}
}

func (builder *SecurityMapperBuilder) UseStrict() *SecurityMapperBuilder { 
  builder.isStrict = true
  return builder
}