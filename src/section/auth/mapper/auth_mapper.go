package auth_mapper

//go:generate mockgen -source=auth_mapper.go -destination=..\..\..\mock\auth\mapper\auth_mapper_mock.go -package=auth_mapper_mock
type AuthMapper interface {
	// DEFINE METHODS
}