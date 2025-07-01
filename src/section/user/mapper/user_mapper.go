package mapper

//go:generate mockgen -source=mapper.go -destination=..\..\..\mock\user\mapper\mapper_mock.go -package=mapper_mock
type UserMapper interface {
	// DEFINE METHODS
}