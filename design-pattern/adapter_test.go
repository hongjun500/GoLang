// @author hongjun500
// @date 2023/10/13 14:41
// @tool ThinkPadX1隐士
// Created with 2022.2.Goland
// Description:

package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestClientAdapter_CreateServer(t *testing.T) {
	var ic ICreateServer

	ic = &AwsClientAdapter{
		Client: AWSClient{},
	}
	err := ic.CreateServer(1, 2)
	assert.NoError(t, err)

	ic = &AliyunClientAdapter{
		Client: AliyunClient{},
	}
	err = ic.CreateServer(3, 4)
	assert.NoError(t, err)

}
