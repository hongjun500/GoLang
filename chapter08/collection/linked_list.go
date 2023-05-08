// Package collection 存放集合的包，ps:利用泛型简单仿写Java中的集合
package collection

import "fmt"

// List 单向链表
type List[T any] struct {
	Next *List[T] // 下一个节点
	Val  T        // 值
}

// Collection 接口包含添加、获取元素方法
type Collection interface {
	AddElement(E List[any]) bool //
	GetElement(index int64) (List[any], error)
}

// AddElement 添加元素
func (l *List[T]) AddElement(E List[any]) bool {
	if l.Next == nil {
		l.Next = &List[T]{Val: E.Val}
	} else {
		l.Next.AddElement(E)
	}
	return false
}

func (l *List[T]) GetElement(index int64) (List[any], error) {
	if index < 0 {
		return List[any]{}, fmt.Errorf("index shoult gt zero")
	}
	return List[any]{}, nil
}
