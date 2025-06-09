package credential_service

//go:generate mockgen -source=credential_service.go -destination=..\..\..\mock\credential\service\credential_service_mock.go -package=credential_service_mock
type CredentialService interface {
	// DEFINE METHODS
}