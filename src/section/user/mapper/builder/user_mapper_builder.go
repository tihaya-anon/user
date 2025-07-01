package builder

import (
	"MVC_DI/section/user/mapper"
	"MVC_DI/section/user/mapper/impl"
	
	"gorm.io/gorm"
)

func (builder *UserMapperBuilder) Build() mapper.UserMapper {
	return builder.userMapperImpl
}

func (builder *UserMapperBuilder) WithDB(DB *gorm.DB) *UserMapperBuilder {
  builder.userMapperImpl.DB = DB
  return builder
}

// BUILDER
type UserMapperBuilder struct {
  isStrict bool
	userMapperImpl *impl.UserMapperImpl
}

func NewUserMapperBuilder() *UserMapperBuilder {
	return &UserMapperBuilder{
		userMapperImpl: &impl.UserMapperImpl{},
	}
}

func (builder *UserMapperBuilder) UseStrict() *UserMapperBuilder { 
  builder.isStrict = true
  return builder
}