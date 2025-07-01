package mapper

//go:generate mockgen -source=mapper.go -destination=..\..\..\mock\permission\mapper\mapper_mock.go -package=mapper_mock
type PermissionMapper interface {
	// DEFINE METHODS
}