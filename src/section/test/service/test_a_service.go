package test_service

//go:generate mockgen -source=test_a_service.go -destination=..\..\..\mock\test\service\test_a_service_mock.go -package=test_service_mock
type TestAService interface {
	// DEFINE METHODS
}