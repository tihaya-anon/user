package auth_mapper_builder

import (
	auth_mapper "MVC_DI/section/auth/mapper"
	auth_mapper_impl "MVC_DI/section/auth/mapper/impl"
	
	"gorm.io/gorm"
)

func (builder *AuthMapperBuilder) Build() auth_mapper.AuthMapper {
	return builder.authMapperImpl
}

func (builder *AuthMapperBuilder) WithDB(DB *gorm.DB) *AuthMapperBuilder {
  builder.authMapperImpl.DB = DB
  return builder
}

// BUILDER
type AuthMapperBuilder struct {
  isStrict bool
	authMapperImpl *auth_mapper_impl.AuthMapperImpl
}

func NewAuthMapperBuilder() *AuthMapperBuilder {
	return &AuthMapperBuilder{
		authMapperImpl: &auth_mapper_impl.AuthMapperImpl{},
	}
}

func (builder *AuthMapperBuilder) UseStrict() *AuthMapperBuilder { 
  builder.isStrict = true
  return builder
}