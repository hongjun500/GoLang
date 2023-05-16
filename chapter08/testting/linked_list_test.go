package testting

import (
	"fmt"
	"github.com/hongjun500/GoLang-master/chapter08/collection"
	"log"
	"testing"
)

func TestLinkedList(t *testing.T) {
	/*javaList := collection.List[string]{
		Val:  "Java",
		Next: nil,
	}*/
	/*	javaList := collection.List[string]{}

		javaList.AddElement("Java")
		javaList.AddElement("Go")
		log.Print(javaList)*/
	list := collection.NewList[string]()
	java := "Java"
	Go := "Go"
	list.AddElement(java)
	list.AddElement(Go)
	log.Print("fsdf", list)

	// 另一种方式，使用 Collection 接口调用 AddElement 方法
	var col collection.Collection[string] = list
	col.AddElement("4")

	// 再次获取列表长度和元素值
	fmt.Printf("List length: %d\n", list.Len())
	fmt.Printf("Element at index 3: %v\n", list.Get(3))
}
