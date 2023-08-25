// @author hongjun500
// @date 2023/8/24 10:59
// @tool ThinkPadX1隐士
// Created with 2022.2.Goland
// Description:

package main

import (
	"container/list"
	"fmt"
	"reflect"
)

// getElem 通过反射可以访问私有字段，但是不能修改私有字段的值
func getElem() {

	/*type List struct {
		root Element // sentinel list element, only &root, root.prev, and root.next are used
		len  int     // current list length excluding (this) sentinel element
	}*/

	l := list.New()
	l.PushFront("foo")
	l.PushFront("bar")

	fv := reflect.ValueOf(l).Elem().FieldByName("len")
	fmt.Printf("list len: %d\n", fv.Int())
	// fv.Set(reflect.ValueOf(100))
}
