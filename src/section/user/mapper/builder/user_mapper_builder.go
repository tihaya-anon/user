package user_mapper_builder

import (
	user_mapper "MVC_DI/section/user/mapper"
	user_mapper_impl "MVC_DI/section/user/mapper/impl"
	
	"gorm.io/gorm"
)

func (builder *UserMapperBuilder) Build() user_mapper.UserMapper {
	return builder.userMapperImpl
}

func (builder *UserMapperBuilder) WithDB(DB *gorm.DB) *UserMapperBuilder {
  builder.userMapperImpl.DB = DB
  return builder
}

// BUILDER
type UserMapperBuilder struct {
  isStrict bool
	userMapperImpl *user_mapper_impl.UserMapperImpl
}

func NewUserMapperBuilder() *UserMapperBuilder {
	return &UserMapperBuilder{
		userMapperImpl: &user_mapper_impl.UserMapperImpl{},
	}
}

func (builder *UserMapperBuilder) UseStrict() *UserMapperBuilder { 
  builder.isStrict = true
  return builder
}