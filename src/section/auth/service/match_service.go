package auth_service

//go:generate mockgen -source=match_service.go -destination=..\..\..\mock\auth\service\match_service_mock.go -package=auth_service_mock
type MatchService interface {
	MatchPassword(raw string, encoded string) bool
	MatchEmailCode(raw string, encoded string) bool
	MatchGoogle2FA(raw string, encoded string) bool
	MatchOauth(raw string, encoded string) bool
}
