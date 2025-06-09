package user_mapper

//go:generate mockgen -source=user_mapper.go -destination=..\..\..\mock\user\mapper\user_mapper_mock.go -package=user_mapper_mock
type UserMapper interface {
	// DEFINE METHODS
}