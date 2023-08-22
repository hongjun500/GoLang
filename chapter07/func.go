// @author hongjun500
// @date 2023/8/22 10:40
// @tool ThinkPadX1隐士
// Created with 2022.2.Goland
// Description:

package main

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

// Info 接收者为值 Teacher
// 由于接收者为值 Teacher, 此方法类的改变属性行为只有引用数据类型 WorkAddress 会被改变，其它的基本数据不会改变
func (t Teacher) Info() {
	// 改变属性 Name 和 Feature 还有 WorkAddress
	// fmt.Println("Info方法，改变属性 Name 和 Feature 还有 WorkAddress 的值")
	t.Name = "改变了姓名的" + t.Name
	t.Age = 2 * t.Age
	t.Sex = 2 * t.Sex
	if t.HasDoWork {
		t.HasDoWork = false
	} else {
		t.HasDoWork = true
	}
	t.Feature = []string{"听歌", "看书"}
	if t.WorkAddress == nil {
		t.WorkAddress = make(map[int]Address)
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

	t.WorkDay = map[int]string{
		1: "两节课" + "1",
		3: "三节课" + "3",
		5: "一节课" + "5",
	}
	// fmt.Printf("teacher.Info 方法执行完毕 %v", t)
}

// Info 接收者为指针 *Student
// 由于接收者为指针 *Student, 此方法内的改变会对其对应原有的属性改变
func (s *Student) Info() {
	s.Name = "改变了姓名的" + s.Name
	s.Age = 223 * 2
	s.Sex = 1000 * 2
	s.HasDoWork = true
	s.Feature = []string{"听歌", "看书"}
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
