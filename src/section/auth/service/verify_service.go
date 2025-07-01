package service

import (
	"MVC_DI/gen/proto"
	"MVC_DI/global/model"
	"MVC_DI/section/auth/dto"
)

type VerifyService interface {
	Verify(dto dto.UserLoginDto, credential *proto.AuthCredential) (bool, proto.LoginResult, *model.AppError)
}
