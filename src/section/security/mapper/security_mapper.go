package mapper

//go:generate mockgen -source=mapper.go -destination=..\..\..\mock\security\mapper\mapper_mock.go -package=mapper_mock
type SecurityMapper interface {
	// DEFINE METHODS
}