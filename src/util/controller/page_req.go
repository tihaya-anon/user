package controller_uitl

import (
	"MVC_DI/vo/req"
	"MVC_DI/vo/resp/common"

	"github.com/gin-gonic/gin"
)

func BindPageReq(ctx *gin.Context) (*req.TPageReq, *common.ValidationError) {
	return BindValidation[req.TPageReq](ctx)
}
