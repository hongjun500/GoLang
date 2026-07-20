package main

import (
	"fmt"
	"reflect"
)

// User 演示结构体：字段上的反引导内容就是 struct tag
type User struct {
	Name  string `json:"name" validate:"required"`
	Age   int    `json:"age"`
	email string // 未导出字段，用来演示 IsExported()
}

func reflectStruct() {
	u := User{Name: "Alice", Age: 30, email: "a@b.com"}
	t := reflect.TypeOf(u)
	v := reflect.ValueOf(u)

	// 遍历结构体字段
	// for i := 0; i < t.NumField(); i++ {
	for i := range t.NumField() { // Go 1.22 可以直接对 int 做 range
		field := t.Field(i) // reflect.StructField: 字段名、类型、tag 等元信息
		value := v.Field(i) // reflect.Value: 字段的实际值
		
		jsonTag := field.Tag.Get("json")
		validateTag := field.Tag.Get("validate")
		fmt.Printf("字段=%-5s 类型=%-6s 导出=%-5v json=%q validate=%q\n",
			field.Name, field.Type, field.IsExported(), jsonTag, validateTag)

		// 未导出字段不能调用 Interface()（会 panic），所以先判断
		if field.IsExported() {
			fmt.Printf("  实际值=%v\n", value.Interface())
		}
	}
}
