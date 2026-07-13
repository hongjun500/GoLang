package main

import (
	"fmt"
	"log"
)

func main() {
	log.Println("自定义类型------")
	// 声明一个名为 strList 类型是 自定义类型 CustomSlice 的变量
	// 基于 []string 的新类型，与 []string 是不同的类型，不能直接赋值
	var strList CustomSlice

	strList = CustomSlice{"1", "2"}
	fmt.Println("变量strList的值为", strList)
	fmt.Printf("变量 strList 的数据类型为%T\n", strList)

	// 自定义类型与底层类型不能直接赋值，需要显式转换
	var raw []string = []string{"a", "b"}
	// 类似 int(i) 这种类型转换，CustomSlice(raw) 是将底层类型为 []string 的 raw 转换为自定义类型 CustomSlice
	strList = CustomSlice(raw)
	fmt.Println("转换后的 strList:", strList)

	log.Println("结构体零值-----------")
	var student struct {
		Name    string
		Version float64
	}
	// 未赋值时各字段为零值："" 和 0
	fmt.Printf("student 零值: %#v\n", student)
	student.Name = "GoLang"
	student.Version = 1.26
	fmt.Printf("student 赋值后: %#v\n", student)


	log.Println("结构体字面量与字段赋值-----------")
	vertex := Vertex{1, 9}
	fmt.Println("vertex:", vertex)

	// 声明一个 Vertex 类型且名称为 vertexData 的变量（可以理解为创建 java 里面的一个对象）
	var vertexData Vertex
	// 设置变量 vertexData 的字段 Y 的值为 1
	vertexData.Y = 1
	fmt.Printf("vertexData: %#v\n", vertexData) // X 仍为 0

	log.Println("自定义类型 vs 类型别名-----------")
	var aint intAlias = 100
	// aint 使用了 类型别名 intAlias，而 intAlias 又是自定义类型 CustomInt, 所以 aint 的数据类型是 main.CustomInt
	fmt.Printf("aint（类型别名）的数据类型是%T \n", aint)

	var cint CustomInt = 100
	fmt.Printf("cint（自定义类型）的数据类型是 %T\n", cint)

	structExmaple()
}
