// @author hongjun500
// @date 2023/10/13 14:20
// @tool ThinkPadX1隐士
// Created with 2022.2.Goland
// Description: 适配器模式

package main

import "fmt"

// ICreateServer 创建云服务器
type ICreateServer interface {
	CreateServer(cpu, mem float64) error
}

// AWSClient aws client
type AWSClient struct {
}

// RunInstance 启动示例
func (c *AWSClient) RunInstance(cpu, mem float64) error {
	fmt.Printf("aws run instance, cpu: %f, mem: %f\n", cpu, mem)
	return nil
}

// AwsClientAdapter aws 适配器
type AwsClientAdapter struct {
	Client AWSClient
}

func (a *AwsClientAdapter) CreateServer(cpu, mem float64) error {
	return a.Client.RunInstance(cpu, mem)
}

// AliyunClient 阿里云客户端
type AliyunClient struct {
}

func (c *AliyunClient) CreateServer(cpu, mem int) error {
	fmt.Printf("aliyun run instance, cpu: %d, mem: %d\n", cpu, mem)
	return nil
}

type AliyunClientAdapter struct {
	Client AliyunClient
}

func (a *AliyunClientAdapter) CreateServer(cpu, mem float64) error {
	return a.Client.CreateServer(int(cpu), int(mem))
}
