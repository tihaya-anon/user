package builder

import (
  "MVC_DI/section/security/service"
  "MVC_DI/section/security/controller"
)

func (builder *SecurityControllerBuilder) Build() *controller.SecurityController {
  if builder.isStrict && builder.securityController.SecurityService == nil {
    panic("`SecurityService` is required")
  }
  return builder.securityController
}

func (builder *SecurityControllerBuilder) WithSecurityService(securityService service.SecurityService) *SecurityControllerBuilder {
  builder.securityController.SecurityService = securityService
  return builder
}

// BUILDER
type SecurityControllerBuilder struct {
  isStrict bool
  securityController *controller.SecurityController
}

func NewSecurityControllerBuilder() *SecurityControllerBuilder {
  return &SecurityControllerBuilder{
    isStrict: false,
    securityController: &controller.SecurityController{},
  }
}

func (builder *SecurityControllerBuilder) UseStrict() *SecurityControllerBuilder { 
  builder.isStrict = true
  return builder
}