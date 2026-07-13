package main

import (
	"fmt"
	"log"
)

func structExmaple() {
	log.Println("匿名结构体-----------")
	// 匿名结构体
	var user struct {
		username string
		password string
	}
	user.username = "admin"
	user.password = "1"
	fmt.Printf("匿名结构体 user %#v \n", user)

	person := struct {
		name string
		age  int
	}{
		name: "john",
		age:  18,
	}
	fmt.Printf("字面量匿名结构体 person %#v\n", person)

	// ------------------- 结构体嵌套（匿名字段提升） -------------------
	circle := Circle{
		Point:  Point{X: 10, Y: 10},
		Radius: 5,
	}
	// 可通过 circle.X 直接访问嵌入的 Point.X
	fmt.Printf("circle.X=%d circle.Y=%d circle.Radius=%d\n", circle.X, circle.Y, circle.Radius)
	fmt.Printf("circle=%#v\n", circle)


	// ------------------- 导出字段 vs 未导出字段 -------------------
	u := User{Name: "Tom"}
	// u.age / u.sex 仅包内可访问；跨包时只有 Name 可见
	u.age = 18
	u.sex = true
	fmt.Printf("User %#v\n", u)
}
