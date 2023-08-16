// @author hongjun500
// @date 2023/8/15 16:32
// @tool ThinkPadX1隐士
// Created with 2022.2.Goland
// Description:

package main

type User struct {
	Name string
	Age  int
}

// NewUser returns a new User
func NewUser(t string) *User {
	switch {
	case "var" == t:
		// 这里的只是声明了一个类型为 User 的指针变量 user，但是并没有初始化，后续如果直接使用 user 会报错
		var user *User
		return user
	case "new" == t:
		return new(User)
	default:
		return &User{}
	}
}
