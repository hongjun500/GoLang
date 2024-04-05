// @author hongjun500
// @date 2024-04-05-22:47
// @tool ThinkPadX1 隐士
// Created with GoLand 2022.2
// Description: common_error_test.go

package tests

import (
	"github.com/hongjun500/GoLang-master/common-web/internal/errors"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCommonErro(t *testing.T) {
	e := &errors.CommonError{}
	t.Logf("e = %v", e)
	assert.Equal(t, e.ErrCode, 0)
	assert.Equal(t, e.Error(), "test error")
	t.Logf("e = %v", e)
}
