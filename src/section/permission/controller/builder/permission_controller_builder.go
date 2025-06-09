package permission_controller_builder

import (
  permission_service "MVC_DI/section/permission/service"
  permission_controller "MVC_DI/section/permission/controller"
)

func (builder *PermissionControllerBuilder) Build() *permission_controller.PermissionController {
  if builder.isStrict && builder.permissionController.PermissionService == nil {
    panic("`PermissionService` is required")
  }
  return builder.permissionController
}

func (builder *PermissionControllerBuilder) WithPermissionService(permissionService permission_service.PermissionService) *PermissionControllerBuilder {
  builder.permissionController.PermissionService = permissionService
  return builder
}

// BUILDER
type PermissionControllerBuilder struct {
  isStrict bool
  permissionController *permission_controller.PermissionController
}

func NewPermissionControllerBuilder() *PermissionControllerBuilder {
  return &PermissionControllerBuilder{
    isStrict: false,
    permissionController: &permission_controller.PermissionController{},
  }
}

func (builder *PermissionControllerBuilder) UseStrict() *PermissionControllerBuilder { 
  builder.isStrict = true
  return builder
}