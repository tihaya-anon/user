package permission_service

//go:generate mockgen -source=permission_service.go -destination=..\..\..\mock\permission\service\permission_service_mock.go -package=permission_service_mock
type PermissionService interface {
	// DEFINE METHODS
}