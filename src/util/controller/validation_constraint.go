package controller_util

import (
	"MVC_DI/util"
	"MVC_DI/vo/resp"
	"MVC_DI/vo/resp/common"
	"errors"
	"net/http"
	"reflect"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func BindValidation[T any](ctx *gin.Context) (*T, *common.ValidationError) {
	var bind T
	err := ctx.ShouldBind(&bind)
	if err != nil {
		var validationErrors validator.ValidationErrors
		validationError := make(common.ValidationError)
		if !errors.As(err, &validationErrors) {
			validationError["error"] = "cannot parse body"
			ctx.AbortWithStatusJSON(http.StatusBadRequest, resp.NewResponse().ValidationError(&validationError))
			return nil, &validationError
		}
		for _, fieldError := range validationErrors {
			field, msg := getValidationMsg(fieldError, &bind)
			if field == "" || msg == "" {
				continue
			}
			validationError[field] = msg
		}
		ctx.AbortWithStatusJSON(http.StatusBadRequest, resp.NewResponse().ValidationError(&validationError))
		return nil, &validationError
	}
	return &bind, nil
}

func getValidationMsg(fieldError validator.FieldError, bind any) (string, string) {
	obj := reflect.TypeOf(bind)
	if field, ok := obj.Elem().FieldByName(fieldError.Field()); ok {
		return util.PascalToSnake(field.Name), field.Tag.Get("msg")
	}
	return "", ""
}
