package handler

import (
	"MVC_DI/global/enum"
	"MVC_DI/global/model"
	"fmt"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func HandleGrpcError(err error) *model.AppError {
	s, ok := status.FromError(err)
	if !ok {
		return model.NewAppError().
			WithStatusKey(enum.SYSTEM_ERROR{}).
			WithDetail(fmt.Sprintf("non-gRPC error: %v", err))
	}

	switch s.Code() {
	case codes.Unavailable, codes.DeadlineExceeded:
		return model.NewAppError().
			WithStatusKey(enum.SERVICE_UNAVAILABLE{}).
			WithDetail(s.Message())

	default:
		return model.NewAppError().
			WithStatusKey(enum.GRPC_ERROR{}).
			WithDetail(fmt.Sprintf("grpc [%s]: %s", s.Code().String(), s.Message()))
	}
}
