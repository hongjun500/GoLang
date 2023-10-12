// @author hongjun500
// @date 2023/10/12 16:46
// @tool ThinkPadX1隐士
// Created with 2022.2.Goland
// Description:

package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestErrorNotification_Notify(t *testing.T) {
	emails := []string{"test@qq.com", "test@gmail.com"}
	emailSender := NewEmailMsgSender(emails)
	notification := NewErrorNotification(emailSender)
	err := notification.Notify("test")
	assert.Nil(t, err)
}
