package mapper

//go:generate mockgen -source=mapper.go -destination=..\..\..\mock\credential\mapper\mapper_mock.go -package=mapper_mock
type CredentialMapper interface {
	// DEFINE METHODS
}