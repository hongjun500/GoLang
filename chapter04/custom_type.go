// @author hongjun500
// @date 2023/8/16 15:03
// @tool ThinkPadX1隐士
// Created with 2022.2.Goland
// Description:

package main

// 自定义类型

// Vertex 定义一个名为 Vertex 的结构体(也代表自定义类型是Vertex)
type Vertex struct {
	X int
	Y int
}
// User 演示导出字段与未导出字段
// Name 首字母大写：可被外部包访问
// age、sex 首字母小写：仅当前包可访问
type User struct {
	Name string
	age  int
	sex  bool
}

type CustomString string

type CustomInt int

type CustomInterface interface {
}

type CustomStruct struct {
}

type CustomFunc func(name string) CustomString

type CustomMap map[string]string

type CustomSlice []string

type CustomInChan chan CustomInt

type CustomPointer *CustomInt

type Point struct {
	X, Y int
}
type Circle struct {
	Point
	Radius int
}
