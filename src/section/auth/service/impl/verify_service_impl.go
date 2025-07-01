package auth_service_impl

import (
	"MVC_DI/gen/proto"
	global_model "MVC_DI/global/model"
	auth_dto "MVC_DI/section/auth/dto"
	auth_enum "MVC_DI/section/auth/enum"
	auth_service "MVC_DI/section/auth/service"
)

type VerifyServiceImpl struct {
	MatchService auth_service.MatchService
}

// Verify implements auth_service.VerifyService.
func (v *VerifyServiceImpl) Verify(dto auth_dto.UserLoginDto, credential *proto.AuthCredential) (bool, proto.LoginResult, *global_model.AppError) {
	var (
		ok     bool
		result proto.LoginResult
		err    *global_model.AppError = global_model.NewAppError()
	)
	switch credential.Type {
	case proto.CredentialType_PASSWORD:
		ok = v.MatchService.MatchPassword(dto.Identifier, dto.Secret, credential.GetSecret())
		result = proto.LoginResult_FAIL_PASSWORD
		err.WithStatusKey(auth_enum.PASSWORD_WRONG{})
	case proto.CredentialType_EMAIL_CODE:
		ok = v.MatchService.MatchEmailCode(dto.Identifier, dto.Secret, credential.GetSecret())
		result = proto.LoginResult_FAIL_EMAIL_CODE
		err.WithStatusKey(auth_enum.EMAIL_CODE_WRONG{})
	case proto.CredentialType__2FA:
		ok = v.MatchService.MatchGoogle2FA(dto.Identifier, dto.Secret, credential.GetSecret())
		result = proto.LoginResult_FAIL_2FA
		err.WithStatusKey(auth_enum.GOOGLE_2FA_WRONG{})
	case proto.CredentialType_OAUTH:
		ok = v.MatchService.MatchOauth(dto.Identifier, dto.Secret, credential.GetSecret())
		result = proto.LoginResult_FAIL_OAUTH
		err.WithStatusKey(auth_enum.OAUTH_WRONG{})
	default:
		ok = false
		result = proto.LoginResult_LOGIN_RESULT_UNSPECIFIED
		err.WithStatusKey(auth_enum.UNKNOWN_CREDENTIAL{})
	}
	if ok {
		result = proto.LoginResult_SUCCESS
		err = nil
	}
	return ok, result, err
}

// INTERFACE
var _ auth_service.VerifyService = (*VerifyServiceImpl)(nil)
