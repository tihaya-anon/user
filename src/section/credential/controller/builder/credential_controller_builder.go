package credential_controller_builder

import (
  credential_service "MVC_DI/section/credential/service"
  credential_controller "MVC_DI/section/credential/controller"
)

func (builder *CredentialControllerBuilder) Build() *credential_controller.CredentialController {
  if builder.isStrict && builder.credentialController.CredentialService == nil {
    panic("`CredentialService` is required")
  }
  return builder.credentialController
}

func (builder *CredentialControllerBuilder) WithCredentialService(credentialService credential_service.CredentialService) *CredentialControllerBuilder {
  builder.credentialController.CredentialService = credentialService
  return builder
}

// BUILDER
type CredentialControllerBuilder struct {
  isStrict bool
  credentialController *credential_controller.CredentialController
}

func NewCredentialControllerBuilder() *CredentialControllerBuilder {
  return &CredentialControllerBuilder{
    isStrict: false,
    credentialController: &credential_controller.CredentialController{},
  }
}

func (builder *CredentialControllerBuilder) UseStrict() *CredentialControllerBuilder { 
  builder.isStrict = true
  return builder
}