package credential_mapper

//go:generate mockgen -source=credential_mapper.go -destination=..\..\..\mock\credential\mapper\credential_mapper_mock.go -package=credential_mapper_mock
type CredentialMapper interface {
	// DEFINE METHODS
}
