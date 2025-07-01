package service

//go:generate mockgen -source=security_service.go -destination=..\..\..\mock\security\service\security_service_mock.go -package=service_mock
type SecurityService interface {
	// DEFINE METHODS
}