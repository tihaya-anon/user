package security_mapper

//go:generate mockgen -source=security_mapper.go -destination=..\..\..\mock\security\mapper\security_mapper_mock.go -package=security_mapper_mock
type SecurityMapper interface {
	// DEFINE METHODS
}