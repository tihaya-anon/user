package impl

import (
	"MVC_DI/gen/proto"
	"MVC_DI/global/model"
	"MVC_DI/section/auth/dto"
	"MVC_DI/section/auth/enum"
	"MVC_DI/section/auth/service"
)

type VerifyServiceImpl struct {
	MatchService service.MatchService
}

// Verify implements service.VerifyService.
func (v *VerifyServiceImpl) Verify(dto dto.UserLoginDto, credential *proto.AuthCredential) (bool, proto.LoginResult, *model.AppError) {
	var (
		ok     bool
		result proto.LoginResult
		err    *model.AppError = model.NewAppError()
	)
	switch credential.Type {
	case proto.CredentialType_PASSWORD:
		ok = v.MatchService.MatchPassword(dto.Identifier, dto.Secret, credential.GetSecret())
		result = proto.LoginResult_FAIL_PASSWORD
		err.WithStatusKey(enum.PASSWORD_WRONG{})
	case proto.CredentialType_EMAIL_CODE:
		ok = v.MatchService.MatchEmailCode(dto.Identifier, dto.Secret, credential.GetSecret())
		result = proto.LoginResult_FAIL_EMAIL_CODE
		err.WithStatusKey(enum.EMAIL_CODE_WRONG{})
	case proto.CredentialType__2FA:
		ok = v.MatchService.MatchGoogle2FA(dto.Identifier, dto.Secret, credential.GetSecret())
		result = proto.LoginResult_FAIL_2FA
		err.WithStatusKey(enum.GOOGLE_2FA_WRONG{})
	case proto.CredentialType_OAUTH:
		ok = v.MatchService.MatchOauth(dto.Identifier, dto.Secret, credential.GetSecret())
		result = proto.LoginResult_FAIL_OAUTH
		err.WithStatusKey(enum.OAUTH_WRONG{})
	default:
		ok = false
		result = proto.LoginResult_LOGIN_RESULT_UNSPECIFIED
		err.WithStatusKey(enum.UNKNOWN_CREDENTIAL{})
	}
	if ok {
		result = proto.LoginResult_SUCCESS
		err = nil
	}
	return ok, result, err
}

// INTERFACE
var _ service.VerifyService = (*VerifyServiceImpl)(nil)
