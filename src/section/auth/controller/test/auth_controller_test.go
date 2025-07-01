package auth_controller_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"MVC_DI/gen/api"
	"MVC_DI/global/enum"
	auth_service_mock "MVC_DI/mock/auth/service"
	auth_controller "MVC_DI/section/auth/controller"
	auth_dto "MVC_DI/section/auth/dto"
	"MVC_DI/vo/resp"

	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
)

func setupRouter() *gin.Engine {
	gin.SetMode(gin.TestMode)
	return gin.Default()
}

func TestLoginUser_Success(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockService := auth_service_mock.NewMockAuthService(ctrl)
	logger := logrus.New()
	controller := auth_controller.AuthController{AuthService: mockService, Logger: logger}

	r := setupRouter()
	r.POST("/login", func(ctx *gin.Context) {
		resp := controller.LoginUser(ctx)
		ctx.JSON(http.StatusOK, resp)
	})

	loginReq := api.UserLoginRequest{
		Secret:     "test_password",
		Identifier: "user@example.com",
		Type:       "email",
	}

	loginResp := &auth_dto.UserLoginRespDto{
		SessionId: 12345,
		Token:     "mock_token",
	}

	mockService.EXPECT().LoginUser(gomock.Any(), gomock.Any()).Return(loginResp, nil)

	body, _ := json.Marshal(loginReq)
	req := httptest.NewRequest(http.MethodPost, "/login", bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	var response resp.TResponse
	_ = json.Unmarshal(w.Body.Bytes(), &response)
	assert.Equal(t, enum.CODE_SUCCESS, response.Code)
	assert.NotNil(t, response.Data)
}

func TestLogoutUser_Success(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockService := auth_service_mock.NewMockAuthService(ctrl)
	logger := logrus.New()
	controller := auth_controller.AuthController{AuthService: mockService, Logger: logger}

	r := setupRouter()
	r.POST("/logout", func(ctx *gin.Context) {
		r := controller.LogoutUser(ctx)
		ctx.JSON(http.StatusOK, r)
	})

	mockService.EXPECT().LogoutUser(gomock.Any(), int64(12345)).Return(nil)

	req := httptest.NewRequest(http.MethodPost, "/logout", nil)
	req.AddCookie(&http.Cookie{Name: "sessionId", Value: "12345"})
	w := httptest.NewRecorder()

	r.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)

	var response resp.TResponse
	_ = json.Unmarshal(w.Body.Bytes(), &response)
	assert.Equal(t, enum.CODE_SUCCESS, response.Code)
}

func TestLogoutUser_MissingToken(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockService := auth_service_mock.NewMockAuthService(ctrl)
	logger := logrus.New()
	controller := auth_controller.AuthController{AuthService: mockService, Logger: logger}

	r := setupRouter()
	r.POST("/logout", func(ctx *gin.Context) {
		r := controller.LogoutUser(ctx)
		ctx.JSON(http.StatusOK, r)
	})

	req := httptest.NewRequest(http.MethodPost, "/logout", nil)
	w := httptest.NewRecorder()

	r.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)

	var response resp.TResponse
	_ = json.Unmarshal(w.Body.Bytes(), &response)
	assert.Equal(t, enum.CODE_MISSING_TOKEN, response.Code)
	assert.Equal(t, enum.MSG_MISSING_TOKEN, response.Msg)
}
