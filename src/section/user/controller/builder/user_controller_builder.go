package builder

import (
  "MVC_DI/section/user/service"
  "MVC_DI/section/user/controller"
)

func (builder *UserControllerBuilder) Build() *controller.UserController {
  if builder.isStrict && builder.userController.UserService == nil {
    panic("`UserService` is required")
  }
  return builder.userController
}

func (builder *UserControllerBuilder) WithUserService(userService service.UserService) *UserControllerBuilder {
  builder.userController.UserService = userService
  return builder
}

// BUILDER
type UserControllerBuilder struct {
  isStrict bool
  userController *controller.UserController
}

func NewUserControllerBuilder() *UserControllerBuilder {
  return &UserControllerBuilder{
    isStrict: false,
    userController: &controller.UserController{},
  }
}

func (builder *UserControllerBuilder) UseStrict() *UserControllerBuilder { 
  builder.isStrict = true
  return builder
}