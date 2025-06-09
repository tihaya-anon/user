package test_mapper_builder

import (
	test_mapper "MVC_DI/section/test/mapper"
	test_mapper_impl "MVC_DI/section/test/mapper/impl"
	
	"gorm.io/gorm"
)

func (builder *TestAMapperBuilder) Build() test_mapper.TestAMapper {
	return builder.testAMapperImpl
}

func (builder *TestAMapperBuilder) WithDB(DB *gorm.DB) *TestAMapperBuilder {
  builder.testAMapperImpl.DB = DB
  return builder
}

// BUILDER
type TestAMapperBuilder struct {
  isStrict bool
	testAMapperImpl *test_mapper_impl.TestAMapperImpl
}

func NewTestAMapperBuilder() *TestAMapperBuilder {
	return &TestAMapperBuilder{
		testAMapperImpl: &test_mapper_impl.TestAMapperImpl{},
	}
}

func (builder *TestAMapperBuilder) UseStrict() *TestAMapperBuilder { 
  builder.isStrict = true
  return builder
}