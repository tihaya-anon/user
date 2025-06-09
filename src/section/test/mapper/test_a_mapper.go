package test_mapper

//go:generate mockgen -source=test_a_mapper.go -destination=..\..\..\mock\test\mapper\test_a_mapper_mock.go -package=test_mapper_mock
type TestAMapper interface {
	// DEFINE METHODS
}