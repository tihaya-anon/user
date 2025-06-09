package auth_service

//go:generate mockgen -source=auth_service.go -destination=..\..\..\mock\auth\service\auth_service_mock.go -package=auth_service_mock
type AuthService interface {
	// DEFINE METHODS
}