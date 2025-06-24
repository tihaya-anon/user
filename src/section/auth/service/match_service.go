package auth_service

//go:generate mockgen -source=match_service.go -destination=..\..\..\mock\auth\service\match_service_mock.go -package=auth_service_mock
type MatchService interface {
	MatchPassword(identifier string, raw string, encoded string) bool
	MatchEmailCode(identifier string, raw string, encoded string) bool
	MatchGoogle2FA(identifier string, raw string, encoded string) bool
	MatchOauth(identifier string, raw string, encoded string) bool
}
