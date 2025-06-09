package security_service

//go:generate mockgen -source=security_service.go -destination=..\..\..\mock\security\service\security_service_mock.go -package=security_service_mock
type SecurityService interface {
	// DEFINE METHODS
}