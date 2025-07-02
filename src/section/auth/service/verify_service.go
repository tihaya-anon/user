package service

import (
	"MVC_DI/gen/proto"
	"MVC_DI/global/model"
	"MVC_DI/section/auth/dto"
)

//go:generate mockgen -source=verify_service.go -destination=..\..\..\mock\auth\service\verify_service_mock.go -package=service_mock
type VerifyService interface {
	Verify(dto dto.UserLoginDto, credential *proto.AuthCredential) (bool, proto.LoginResult, *model.AppError)
}
