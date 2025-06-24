package auth_service_impl

import auth_service "MVC_DI/section/auth/service"

type MatchServiceImpl struct {
}

func (m MatchServiceImpl) MatchPassword(identifier string, raw string, encoded string) bool {
	//TODO implement me
	panic("implement me")
}

func (m MatchServiceImpl) MatchEmailCode(identifier string, raw string, encoded string) bool {
	//TODO implement me
	panic("implement me")
}

func (m MatchServiceImpl) MatchGoogle2FA(identifier string, raw string, encoded string) bool {
	//TODO implement me
	panic("implement me")
}

func (m MatchServiceImpl) MatchOauth(identifier string, raw string, encoded string) bool {
	//TODO implement me
	panic("implement me")
}

// INTERFACE
var _ auth_service.MatchService = (*MatchServiceImpl)(nil)
