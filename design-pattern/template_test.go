// @author hongjun500
// @date 2023/10/17 14:40
// @tool ThinkPadX1隐士
// Created with 2022.2.Goland
// Description:

package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTemplate_SMS_Send(t *testing.T) {
	tel := NewTelecomSms()
	err := tel.Send("test", "123456789")
	assert.NoError(t, err)
}
