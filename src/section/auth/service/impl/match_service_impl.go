package auth_service_impl

import auth_service "MVC_DI/section/auth/service"

type MatchServiceImpl struct {
}

func (m MatchServiceImpl) MatchPassword(raw string, encoded string) bool {
	//TODO implement me
	panic("implement me")
}

func (m MatchServiceImpl) MatchEmailCode(raw string, encoded string) bool {
	//TODO implement me
	panic("implement me")
}

func (m MatchServiceImpl) MatchGoogle2FA(raw string, encoded string) bool {
	//TODO implement me
	panic("implement me")
}

func (m MatchServiceImpl) MatchOauth(raw string, encoded string) bool {
	//TODO implement me
	panic("implement me")
}

// INTERFACE
var _ auth_service.MatchService = (*MatchServiceImpl)(nil)
