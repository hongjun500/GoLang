// @author hongjun500
// @date 2023/10/24 14:14
// @tool ThinkPadX1隐士
// Created with 2022.2.Goland
// Description:

package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUserProxyLogin(t *testing.T) {
	proxy := NewUserProxy(&User{})
	err := proxy.Login("test", "123456")
	assert.NoError(t, err)
}
