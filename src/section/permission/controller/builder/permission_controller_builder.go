package builder

import (
  "MVC_DI/section/permission/service"
  "MVC_DI/section/permission/controller"
)

func (builder *PermissionControllerBuilder) Build() *controller.PermissionController {
  if builder.isStrict && builder.permissionController.PermissionService == nil {
    panic("`PermissionService` is required")
  }
  return builder.permissionController
}

func (builder *PermissionControllerBuilder) WithPermissionService(permissionService service.PermissionService) *PermissionControllerBuilder {
  builder.permissionController.PermissionService = permissionService
  return builder
}

// BUILDER
type PermissionControllerBuilder struct {
  isStrict bool
  permissionController *controller.PermissionController
}

func NewPermissionControllerBuilder() *PermissionControllerBuilder {
  return &PermissionControllerBuilder{
    isStrict: false,
    permissionController: &controller.PermissionController{},
  }
}

func (builder *PermissionControllerBuilder) UseStrict() *PermissionControllerBuilder { 
  builder.isStrict = true
  return builder
}