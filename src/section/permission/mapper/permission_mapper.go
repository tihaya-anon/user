package permission_mapper

//go:generate mockgen -source=permission_mapper.go -destination=..\..\..\mock\permission\mapper\permission_mapper_mock.go -package=permission_mapper_mock
type PermissionMapper interface {
	// DEFINE METHODS
}