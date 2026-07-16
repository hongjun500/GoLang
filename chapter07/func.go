// @author hongjun500
// @date 2023/8/22 10:40
// @tool ThinkPadX1隐士
// Created with 2022.2.Goland
// Description:

package main

import "fmt"

// Teacher 声明老师的结构体
type Teacher struct {
	Name      string
	Age       int
	Sex       int
	HasDoWork bool
	// 特点
	Feature     []string        // 引用类型：切片
	WorkAddress map[int]Address // 引用类型：map
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

// Address 地址；longitude/latitude 未导出，仅本包可见
type Address struct {
	AddressName string
	longitude   float64
	latitude    float64
}

// String 实现 fmt.Stringer 接口：用 %v/%s 打印时会自动调用，让输出更可读，
// 同时演示"标准库接口的隐式实现"。
func (t Teacher) String() string {
	return fmt.Sprintf("Teacher{Name:%q, Age:%d}", t.Name, t.Age)
}
func (s *Student) String() string {
	return fmt.Sprintf("Student{Name:%q, Age:%d}", s.Name, s.Age)
}

// Info 值接收者：t 是整个 Teacher 的副本
// - 对基本字段/切片头/map 头的“重新赋值”只改副本，外部不可见
// - 但通过副本里的 map 头去“写元素”，改的是同一块底层 map，外部可见
func (t Teacher) Info() {
	// 改变属性 Name 和 Feature 还有 WorkAddress
	t.Name = "改变了姓名的" + t.Name // 只改副本，外部不可见
	t.Age = 2 * t.Age          // 只改副本
	t.Sex = 2 * t.Sex          // 只改副本
	if t.HasDoWork {
		t.HasDoWork = false // 只改副本
	} else {
		t.HasDoWork = true // 只改副本
	}
	t.Feature = []string{"听歌", "看书"} // 重新赋值切片头，外部不可见
	if t.WorkAddress == nil {
		t.WorkAddress = make(map[int]Address)
	}
	// 往同一底层 map 写元素，外部可见
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

	t.WorkDay = map[int]string{
		1: "两节课" + "1",
		3: "三节课" + "3",
		5: "一节课" + "5",
	}
}

// Info 指针接收者 *Student
// 由于接收者为指针 *Student, 此方法内的改变会对其对应原有的属性改变
// 指针接收者：s 指向原对象，所有修改都直接作用于原对象，外部全部可见。
func (s *Student) Info() {
	s.Name = "改变了姓名的" + s.Name
	s.Age = 223 * 2
	s.Sex = 1000 * 2
	s.HasDoWork = true
	s.Feature = []string{"听歌", "看书"}
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
	s.WorkDay = map[int]string{
		1: "十二节课" + "1",
		3: "十二节课" + "3",
		5: "十二节课" + "5",
	}
	// fmt.Printf("student.Info 方法执行完毕 %v", *s)
}

// init 在 main 之前自动执行，准备演示数据
func init() {

	teacher = Teacher{
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
	// 注意：Student 用指针接收者实现 Info，因此必须存 *Student
	student = &Student{
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
}
