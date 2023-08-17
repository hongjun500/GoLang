package main

import (
	"fmt"
)

func main() {

	// 声明一个名为 strList 类型是 自定义类型 CustomSlice 的变量
	var strList CustomSlice

	strList = CustomSlice{"1", "2"}
	fmt.Println("变量strList的值为", strList)
	fmt.Printf("变量 strList 的数据类型为%T", strList)
	var student struct {
		Name    string
		Version float64
	}
	// student.Name = "goLang"
	// student.Version = 1.17
	fmt.Println(student.Name)
	fmt.Println(student.Version)

	vertex := Vertex{1, 9}
	fmt.Println(vertex)

	// 声明一个 Vertex 类型且名称为 vertexData 的变量（可以理解为创建 java 里面的一个对象）
	var vertexData Vertex
	// 设置变量 vertexData 的字段 Y 的值为 1
	vertexData.Y = 1

	var aint intAlias = 100
	// aint 使用了 类型别名 intAlias，而 intAlias 又是自定义类型 CustomInt, 所以 aint 的数据类型是 main.CustomInt
	fmt.Printf("aint 的数据类型是%T \n", aint)

	// 匿名结构体
	var user struct {
		username string
		password string
	}
	user.username = "admin"
	user.password = "1"
	fmt.Printf("匿名结构体 user %#v \n", user)

}
