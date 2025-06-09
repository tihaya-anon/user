package test_mapper_impl

import (
	test_mapper "MVC_DI/section/test/mapper"

	"gorm.io/gorm"
)

type TestAMapperImpl struct{
	DB *gorm.DB
}

// INTERFACE
var _ test_mapper.TestAMapper = (*TestAMapperImpl)(nil)