package collection

/**
 * @author hongjun500
 * @date 2023/5/8
 * @tool ThinkPadX1隐士
 * Created with GoLand 2022.2
 * Description: list.go
 */

type List[T any] struct {
	next *List[T]
	val  T
}

func NewList[T any]() *List[T] {
	return &List[T]{}
}

type Collection[T any] interface {
	AddElement(val T)
	Get(index int) T
	Len() int
}

func (l *List[T]) AddElement(val T) {
	if l == nil {
		return
	}
	for l.next != nil {
		l = l.next
	}
	l.next = &List[T]{val: val}
}

func (l *List[T]) Get(index int) T {
	if l == nil {
		var zero T
		return zero
	}
	for i := 0; i < index; i++ {
		if l.next == nil {
			var zero T
			return zero
		}
		l = l.next
	}
	return l.val
}

func (l *List[T]) Len() int {
	if l == nil {
		return 0
	}
	count := 0
	for l.next != nil {
		count++
		l = l.next
	}
	return count + 1
}
