// @author hongjun500
// @date 2023/8/31 14:21
// @tool ThinkPadX1隐士
// Created with 2022.2.Goland
// Description:

package linkedlist

import (
	"testing"
)

// TestRemoveElements 测试移除单向链表元素
func TestRemoveElements(t *testing.T) {
	node1 := &LinkedNode{Val: 1}
	node2 := &LinkedNode{Val: 2}
	node3 := &LinkedNode{Val: 3}
	node4 := &LinkedNode{Val: 4}
	node5 := &LinkedNode{Val: 5}
	node1.Next = node2
	node2.Next = node3
	node3.Next = node4
	node4.Next = node5
	head := removeElements(node1, 3)
	for head != nil {
		t.Log(head.Val)
		head = head.Next
	}
}

func TestMyLinkedList(t *testing.T) {
	cur := Constructor()
	list := &cur
	list.AddAtHead(1)
	list.AddAtTail(3)
	list.AddAtIndex(1, 2)
	t.Log(list.Get(1))

	list.DeleteAtIndex(1)

	t.Log(list.Get(1))

	t.Log(list.Get(3))

	list.DeleteAtIndex(3)

	list.DeleteAtIndex(0)

	t.Log(list.Get(0))
	list.DeleteAtIndex(0)
	t.Log(list.Get(0))
}
