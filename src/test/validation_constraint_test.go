package test

import (
	controller_util "MVC_DI/util/controller"
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

type TestRequest struct {
	Name string `json:"name" binding:"required" msg:"not null"`
	Age  int    `json:"age" binding:"gte=1,lte=120" msg:"must between 1 and 120"`
}

func TestBindValidation(t *testing.T) {
	gin.SetMode(gin.TestMode)

	router := gin.Default()
	router.POST("/test", func(c *gin.Context) {
		data, validationErr := controller_util.BindValidation[TestRequest](c)
		if validationErr != nil {
			t.Log(validationErr)
			c.JSON(http.StatusBadRequest, validationErr)
			return
		}
		c.JSON(http.StatusOK, data)
	})

	t.Run("valid input", func(t *testing.T) {
		body := map[string]any{"name": "Alice", "age": 30}
		jsonBytes, _ := json.Marshal(body)

		req, _ := http.NewRequest("POST", "/test", bytes.NewBuffer(jsonBytes))
		req.Header.Set("Content-Type", "application/json")
		w := httptest.NewRecorder()

		router.ServeHTTP(w, req)
		assert.Equal(t, 200, w.Code)
	})

	t.Run("missing name", func(t *testing.T) {
		body := map[string]any{"age": 30}
		jsonBytes, _ := json.Marshal(body)

		req, _ := http.NewRequest("POST", "/test", bytes.NewBuffer(jsonBytes))
		req.Header.Set("Content-Type", "application/json")
		w := httptest.NewRecorder()

		router.ServeHTTP(w, req)
		assert.Equal(t, 400, w.Code)
		assert.Contains(t, w.Body.String(), "not null")
	})

	t.Run("invalid age", func(t *testing.T) {
		body := map[string]any{"name": "Bob", "age": 130}
		jsonBytes, _ := json.Marshal(body)

		req, _ := http.NewRequest("POST", "/test", bytes.NewBuffer(jsonBytes))
		req.Header.Set("Content-Type", "application/json")
		w := httptest.NewRecorder()

		router.ServeHTTP(w, req)
		assert.Equal(t, 400, w.Code)
		assert.Contains(t, w.Body.String(), "must between 1 and 120")
	})
}
