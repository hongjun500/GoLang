// @author hongjun500
// @date 2023/10/24 14:04
// @tool ThinkPadX1隐士
// Created with 2022.2.Goland
// Description: 代理模式

package main

import (
	"log"
	"time"
)

type IUser interface {
	Login(username, password string) error
}

type User struct {
}

func (User) Login(username, password string) error {
	return nil
}

type UserProxy struct {
	user *User
}

func NewUserProxy(user *User) *UserProxy {
	return &UserProxy{user: user}
}

func (u *UserProxy) Login(username, password string) error {
	start := time.Now()
	if err := u.user.Login(username, password); err != nil {
		return err
	}
	// after 这里可能也有一些监控统计的逻辑
	log.Printf("user login cost time: %s", time.Now().Sub(start))
	return nil
}
