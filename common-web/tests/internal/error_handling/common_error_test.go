// @author hongjun500
// @date 2024-04-05-22:47
// @tool ThinkPadX1 隐士
// Created with GoLand 2022.2
// Description: common_error_test.go

package error_handling

import (
	"errors"
	"strings"
	"testing"

	"github.com/hongjun500/GoLang-master/common-web/internal/error_handling"
	"github.com/stretchr/testify/assert"
)

func TestErrorConstant(t *testing.T) {
	assert.Equal(t, error_handling.UnknownError, 100000)
	assert.Equal(t, error_handling.CustomError, 100001)
	assert.Equal(t, error_handling.ParameterValidationError, 100002)
	assert.Equal(t, error_handling.DatabaseError, 100003)

	msg, ok := error_handling.BusinessErrorConstant[error_handling.UnknownError]
	assert.Equal(t, "未知错误", msg)
	assert.Equal(t, true, ok)
}

var newerr = errors.New("error_handling is not nil")

func TestCommonErrorOneArgument(t *testing.T) {
	// panic
	assert.Panics(t, func() { _ = error_handling.NewCommonError("test error_handling") })

	err := error_handling.NewCommonError(newerr)
	t.Logf("error_handling = %v", err)
	assert.Equal(t, newerr, err.Err)
	assert.Equal(t, "error_handling is not nil", err.Err.Error())

	err = error_handling.NewCommonError(error_handling.UnknownError)
	t.Logf("error_handling = %v", err)
	assert.Nil(t, err.Err, "error_handling.Err is nil")
	assert.Equal(t, error_handling.UnknownError, err.ErrCode)
	assert.Equal(t, "未知错误", err.ErrMsg, err.Error())
}

func TestCommonErrorOneArgumentv2(t *testing.T) {
	// panic
	assert.Panics(t, func() { _ = error_handling.NewCommonErrorV2("test error_handling") })

	err := error_handling.NewCommonErrorV2(newerr)
	t.Logf("error_handling = %v", err)
	assert.Equal(t, newerr, err.Err)
	assert.Equal(t, "error_handling is not nil", err.Err.Error())

	err = error_handling.NewCommonErrorV2(error_handling.UnknownError)
	t.Logf("error_handling = %v", err)
	assert.Nil(t, err.Err, "error_handling.Err is nil")
	assert.Equal(t, error_handling.UnknownError, err.ErrCode)
	assert.Equal(t, "未知错误", err.ErrMsg, err.Error())
}

func TestCommonErrorTwoArgument(t *testing.T) {
	// panic
	assert.Panics(t, func() { _ = error_handling.NewCommonError(newerr, "test error_handling") })

	err := error_handling.NewCommonError(newerr, error_handling.ParameterValidationError)
	t.Logf("error_handling = %v", err)
	assert.Equal(t, newerr, err.Err)
	assert.Equal(t, error_handling.ParameterValidationError, err.ErrCode)
	assert.Equal(t, "参数校验错误", err.ErrMsg, err.Error())

	err = error_handling.NewCommonError(error_handling.UnknownError, "test error_handling")
	t.Logf("error_handling = %v", err)
	assert.Nil(t, nil, err.Err)
	assert.Equal(t, error_handling.UnknownError, err.ErrCode)
	assert.Equal(t, "test error_handling", err.Error(), err.ErrMsg)
}
func TestCommonErrorTwoArgumentV2(t *testing.T) {
	// panic
	assert.Panics(t, func() { _ = error_handling.NewCommonErrorV2(newerr, "test error_handling") })

	err := error_handling.NewCommonErrorV2(newerr, error_handling.ParameterValidationError)
	t.Logf("error_handling = %v", err)
	assert.Equal(t, newerr, err.Err)
	assert.Equal(t, error_handling.ParameterValidationError, err.ErrCode)
	assert.Equal(t, "参数校验错误", err.ErrMsg, err.Error())

	err = error_handling.NewCommonErrorV2(error_handling.UnknownError, "test error_handling")
	t.Logf("error_handling = %v", err)
	assert.Nil(t, nil, err.Err)
	assert.Equal(t, error_handling.UnknownError, err.ErrCode)
	assert.Equal(t, "test error_handling", err.Error(), err.ErrMsg)
}

func TestCommonErrorThreeArgument(t *testing.T) {
	err := error_handling.NewCommonError(newerr, error_handling.DatabaseError, "test database error_handling")
	t.Logf("error_handling = %s", err)
	assert.Equal(t, newerr, err.Err)
	assert.Equal(t, error_handling.DatabaseError, err.ErrCode)
	assert.Equal(t, "test database error_handling", err.ErrMsg, err.Error())
	assert.True(t, true, strings.EqualFold(err.ErrMsg, err.Error()))
	assert.True(t, true, err.ErrMsg == err.Error())
}

func TestCommonErrorThreeArgumentV2(t *testing.T) {
	err := error_handling.NewCommonErrorV2(newerr, error_handling.DatabaseError, "test database error_handling")
	t.Logf("error_handling = %s", err)
	assert.Equal(t, newerr, err.Err)
	assert.Equal(t, error_handling.DatabaseError, err.ErrCode)
	assert.Equal(t, "test database error_handling", err.ErrMsg, err.Error())
	assert.True(t, true, strings.EqualFold(err.ErrMsg, err.Error()))
	assert.True(t, true, err.ErrMsg == err.Error())
}
