package main

import (
	"fmt"
)

// People 定义一个接口
type People interface {
	// 简介的方法
	// Info()
	// 简介的方法，入参是指针
	InfoPointer()
}

// Teacher 声明老师的结构体
type Teacher struct {
	Name      string
	Age       int
	Sex       int
	HasDoWork bool
	// 特点
	Feature     []string
	WorkAddress map[int]Address
	WorkDay     map[int]string
}

// Student 声明学生的结构体
type Student struct {
	Name      string
	Age       int
	Sex       int
	HasDoWork bool
	// 特点
	Feature     []string
	WorkAddress map[int][]string
	WorkDay     map[int]string
}

type Address struct {
	AddressName string
	longitude   float64
	latitude    float64
}

func describe(i int) {
	fmt.Printf("(%v, %T)\n", i, i)
}

// InfoPointer 接收者为值 Teacher
// 由于接收者为值 Teacher, 此方法类的改变属性行为只有引用数据类型WorkAddress会被改变，其它的基本数据不会改变
func (t Teacher) InfoPointer() {
	// 改变属性Name和Feature还有WorkAddress
	// fmt.Println("Info方法，改变属性Name和Feature还有WorkAddress")
	t.Name = "改变了姓名的" + t.Name
	t.Age = 999
	t.Sex = 0
	t.HasDoWork = true
	// fmt.Println("这是" + t.Name + "老师，" + "年龄=" + strconv.Itoa(t.Age) + "岁" +
	// 	", 性别为" + strconv.Itoa(t.Sex) + ",是否有工作？" + strconv.FormatBool(t.HasDoWork))
	t.Feature = []string{"听歌", "看书"}
	// fmt.Printf("特点是: %v \n", t.Feature)
	if t.WorkAddress == nil {
		t.WorkAddress = make(map[int]Address, 0)
	}
	t.WorkAddress[1] = Address{
		AddressName: "游泳池",
		longitude:   11111,
		latitude:    22222,
	}
	t.WorkAddress[2] = Address{
		AddressName: "5A级景区",
		longitude:   11111,
		latitude:    22222,
	}

	// fmt.Printf("工作地点是: %v \n", t.WorkAddress)
	t.WorkDay = map[int]string{
		1: "两节课",
		3: "三节课",
		5: "一节课",
	}
}

// InfoPointer 接收者为指针 *Student
// 由于接收者为指针 *Student, 此方法内的改变会对其对应原有的属性改变
func (s *Student) InfoPointer() {
	// 改变属性Name和Feature还有WorkAddress
	// fmt.Println("InfoPointer方法，改变属性Name和Feature还有WorkAddress")
	s.Name = "改变了姓名的" + s.Name
	s.Age = 223 * 2
	s.Sex = 1000 * 2
	s.HasDoWork = true
	s.Feature = []string{"听歌", "看书"}
	// fmt.Println("这是" + t.Name + "老师，" + "年龄=" + strconv.Itoa(t.Age) + "岁" +
	// 	", 性别为" + strconv.Itoa(t.Sex) + ",是否有工作？" + strconv.FormatBool(t.HasDoWork))
	s.Feature = []string{"听歌", "看书"}
	// fmt.Printf("特点是: %v \n", t.Feature)
	s.WorkAddress = map[int][]string{
		1: {
			"游泳池",
			"11111",
			"22222",
		},
		2: {
			"5A级景区",
			"11111",
			"22222",
		},
	}
	// fmt.Printf("工作地点是: %v \n", t.WorkAddress)
	s.WorkDay = map[int]string{
		1: "十二节课",
		3: "十二节课",
		5: "十二节课",
	}
}

func main() {
	teacherPeople := Teacher{
		Name:      "李青老师",
		Age:       35,
		Sex:       2,
		HasDoWork: false,
		Feature:   []string{"上课", "改作业"},
		WorkAddress: map[int]Address{
			1: {
				AddressName: "502教室",
				longitude:   50000.502,
				latitude:    20000.502,
			},
			2: {
				AddressName: "609教室",
				longitude:   66666.609,
				latitude:    90000.609,
			},
		},
		WorkDay: map[int]string{
			1: "两节课",
			3: "三节课",
			5: "两节课",
		},
	}
	fmt.Println("teacherPeople原属性：")
	fmt.Println(teacherPeople)
	fmt.Println("teacherPeople调用InfoPointer方法得到的数据")
	teacherPeople.InfoPointer()
	t := new(Teacher)
	t.InfoPointer()
	fmt.Println(teacherPeople)
	fmt.Println("------------------------------------------------------------------------")

	studentPeople := Student{
		Name:      "李青学生",
		Age:       18,
		Sex:       1,
		HasDoWork: false,
		Feature:   []string{"上课", "写作业"},
		WorkAddress: map[int][]string{
			1: {
				"502教室",
				"50000.502",
				"20000.502",
			},
			2: {
				"609教室",
				"66666.609",
				"90000.609",
			},
		},
		WorkDay: map[int]string{
			1: "十节课",
			3: "十节课",
			5: "十节课",
		},
	}

	fmt.Println("studentPeople原属性：")
	fmt.Println(studentPeople)
	s := &studentPeople
	s.InfoPointer()
	fmt.Println("studentPeople调用InfoPointer方法得到的数据")
	studentPeople.InfoPointer()
	fmt.Println(studentPeople)

}

// SumIntsOrFloats sums the values of map m. It supports both int64 and float64
// as types for map values.
func SumIntsOrFloats[K comparable, V int64 | float64](m map[K]V) V {
	var s V
	for _, v := range m {
		s += v
	}
	return s
}
