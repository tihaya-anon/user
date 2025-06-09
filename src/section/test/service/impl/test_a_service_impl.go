package test_service_impl

import (
	test_service "MVC_DI/section/test/service"
	test_mapper "MVC_DI/section/test/mapper"
)

type TestAServiceImpl struct{
	TestAMapper test_mapper.TestAMapper
}

// INTERFACE
var _ test_service.TestAService = (*TestAServiceImpl)(nil)
