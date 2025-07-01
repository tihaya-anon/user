package controller_builder

import (
	"MVC_DI/section/auth/controller"
	"MVC_DI/section/auth/service"
)

func (builder *AuthControllerBuilder) Build() *controller.AuthController {
	if builder.isStrict && builder.authController.AuthService == nil {
		panic("`AuthService` is required")
	}
	return builder.authController
}

func (builder *AuthControllerBuilder) WithAuthService(authService service.AuthService) *AuthControllerBuilder {
	builder.authController.AuthService = authService
	return builder
}

// BUILDER
type AuthControllerBuilder struct {
	isStrict       bool
	authController *controller.AuthController
}

func NewAuthControllerBuilder() *AuthControllerBuilder {
	return &AuthControllerBuilder{
		isStrict:       false,
		authController: &controller.AuthController{},
	}
}

func (builder *AuthControllerBuilder) UseStrict() *AuthControllerBuilder {
	builder.isStrict = true
	return builder
}
