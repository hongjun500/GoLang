package main

import "fmt"

// StudentInterface 定义一个接口
type StudentInterface interface {
	PrintAge()
}

type Student struct {
	Name string
	Age  int
}

func (s Student) PrintAge() {
	i := 0
	for i < s.Age {
		fmt.Print("循环Age=", s.Age, "次打印出年龄\n")
		i++
	}
}

func newStudent(age int, name string) StudentInterface {
	return Student{name, age}
}

func main() {
	st := Student{
		Name: "李青",
		Age:  3,
	}
	st.PrintAge()

	student := newStudent(1, "张三")
	student.PrintAge()
}
