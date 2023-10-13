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

func TestAwsClientAdapter_CreateServer(t *testing.T) {
	aws := &AwsClientAdapter{}
	aws.Client = AWSClient{}
	err := aws.CreateServer(1, 2)
	assert.NoError(t, err)
}

func TestAliyunClientAdapter_CreateServer(t *testing.T) {
	aliyun := &AliyunClientAdapter{}
	aliyun.Client = AliyunClient{}
	err := aliyun.CreateServer(3, 4)
	assert.NoError(t, err)
}
