// @author hongjun500
// @date 2024-04-06-21:25
// @tool ThinkPadX1 隐士
// Created with GoLand 2022.2
// Description: common_return_type_test.go

package response

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/hongjun500/GoLang-master/common-web/internal/error_handling"
	"github.com/hongjun500/GoLang-master/common-web/internal/response"
	"github.com/stretchr/testify/assert"
)

var router *gin.Engine

var recorder *httptest.ResponseRecorder

func init() {
	gin.SetMode(gin.TestMode)
	router = gin.Default()

	recorder = httptest.NewRecorder()
}

func TestCreated(t *testing.T) {
	gin.SetMode(gin.TestMode)
	router = gin.Default()

	router.GET("/test", func(context *gin.Context) {
		response.Created(context)
	})

	request, _ := http.NewRequest(http.MethodGet, "/test", nil)
	recorder = httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	assert.Equal(t, http.StatusOK, recorder.Code)
	assert.Contains(t, recorder.Body.String(), "success")
	assert.Contains(t, recorder.Body.String(), "null")
}

func TestCreateSuccessWithData(t *testing.T) {

	var mapData = map[string]any{
		"key":  "value",
		"key2": 123,
		"say":  "hello",
	}

	router.GET("/test", func(context *gin.Context) {
		response.CreateSuccess(context, mapData)
	})

	request, _ := http.NewRequest(http.MethodGet, "/test", nil)

	router.ServeHTTP(recorder, request)

	assert.Equal(t, http.StatusOK, recorder.Code)
	assert.Contains(t, recorder.Body.String(), "success")

	var returnType response.CommonReturnType
	e := json.Unmarshal(recorder.Body.Bytes(), &returnType)
	assert.Nil(t, e)
	m := returnType.Data.(map[string]any)
	assert.Equal(t, "value", m["key"])
	assert.Equal(t, 123, int(m["key2"].(float64)))
	assert.Equal(t, "hello", m["say"])
}

func TestCreateFailWithData(t *testing.T) {

	router.GET("/test", func(context *gin.Context) {
		response.CreateFail(context, error_handling.NewCommonError(error_handling.CustomError))
	})

	request, _ := http.NewRequest(http.MethodGet, "/test", nil)
	recorder = httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	assert.Equal(t, http.StatusOK, recorder.Code)
	assert.Contains(t, recorder.Body.String(), "fail")

	var returnType response.CommonReturnType
	e := json.Unmarshal(recorder.Body.Bytes(), &returnType)
	assert.Nil(t, e)
	m := returnType.Data.(map[string]any)
	// convert float64 to int64
	assert.Equal(t, error_handling.CustomError, int(m["err_code"].(float64)))
	assert.Equal(t, "自定义错误", m["err_msg"])
}
