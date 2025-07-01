package builder

import (
  "MVC_DI/section/credential/service"
  "MVC_DI/section/credential/controller"
)

func (builder *CredentialControllerBuilder) Build() *controller.CredentialController {
  if builder.isStrict && builder.credentialController.CredentialService == nil {
    panic("`CredentialService` is required")
  }
  return builder.credentialController
}

func (builder *CredentialControllerBuilder) WithCredentialService(credentialService service.CredentialService) *CredentialControllerBuilder {
  builder.credentialController.CredentialService = credentialService
  return builder
}

// BUILDER
type CredentialControllerBuilder struct {
  isStrict bool
  credentialController *controller.CredentialController
}

func NewCredentialControllerBuilder() *CredentialControllerBuilder {
  return &CredentialControllerBuilder{
    isStrict: false,
    credentialController: &controller.CredentialController{},
  }
}

func (builder *CredentialControllerBuilder) UseStrict() *CredentialControllerBuilder { 
  builder.isStrict = true
  return builder
}