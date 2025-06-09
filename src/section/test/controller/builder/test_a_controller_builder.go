package test_controller_builder

import (
  test_service "MVC_DI/section/test/service"
  test_controller "MVC_DI/section/test/controller"
)

func (builder *TestAControllerBuilder) Build() *test_controller.TestAController {
  if builder.isStrict && builder.testAController.TestAService == nil {
    panic("`TestAService` is required")
  }
  return builder.testAController
}

func (builder *TestAControllerBuilder) WithTestAService(testAService test_service.TestAService) *TestAControllerBuilder {
  builder.testAController.TestAService = testAService
  return builder
}

// BUILDER
type TestAControllerBuilder struct {
  isStrict bool
  testAController *test_controller.TestAController
}

func NewTestAControllerBuilder() *TestAControllerBuilder {
  return &TestAControllerBuilder{
    isStrict: false,
    testAController: &test_controller.TestAController{},
  }
}

func (builder *TestAControllerBuilder) UseStrict() *TestAControllerBuilder { 
  builder.isStrict = true
  return builder
}