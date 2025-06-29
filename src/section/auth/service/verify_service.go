package auth_service

import (
	"MVC_DI/gen/proto"
	global_model "MVC_DI/global/model"
	auth_dto "MVC_DI/section/auth/dto"
)

type VerifyService interface {
	Verify(dto auth_dto.UserLoginDto, credential *proto.AuthCredential) (bool, proto.LoginResult, *global_model.AppError)
}
