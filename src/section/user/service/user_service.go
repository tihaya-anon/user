package service

//go:generate mockgen -source=user_service.go -destination=..\..\..\mock\user\service\user_service_mock.go -package=service_mock
type UserService interface {
	// DEFINE METHODS
}