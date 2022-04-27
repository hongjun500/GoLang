package main

import (
	"fmt"
)

// 定义一个名为Vertex的结构体(也代表自定义类型是Vertex)
type Vertex struct {
	X int
	Y int
}

// 自定义一个数据类型，数据类型为字符串切片，命名为stringList
type stringList []string

// 类型别名
type aliasInt = int

func main() {

	// 声明一个名为strList类型是stringList的变量
	var strList stringList

	// strList = stringList{"1","2"}
	// 也可以直接
	//strList := stringList{"hello", "go"}
	fmt.Println("变量strList的值为", strList)
	fmt.Printf("变量strList的数据类型为%T", strList)
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

	// 声明一个Vertex类型且名称为vertexData的变量（可以理解为创建java里面的一个对象）
	var vertexData Vertex
	// 设置变量vertexData的字段Y的值为1
	vertexData.Y = 1

	var aint aliasInt = 100
	fmt.Printf("aint的数据类型是%T \n", aint)

	// 匿名结构体
	var user struct {
		username string
		password string
	}
	user.username = "admin"
	user.password = "1"
	fmt.Printf("%#v \n", user)

}
