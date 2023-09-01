// @author hongjun500
// @date 2023/8/31 11:00
// @tool ThinkPadX1隐士
// Created with 2022.2.Goland
// Description:

package linkedlist

type LinkedNode struct {
	Val  int
	Next *LinkedNode
}

// MyLinkedList 单向链表
type MyLinkedList struct {
	LinkedNode *LinkedNode
	Size       int
}

func (list *MyLinkedList) AddAtIndex(index int, val int) {
	if index == list.Size {
		list.AddAtTail(val)
		return
	}
	if index < list.Size {
		cur := list.LinkedNode
		for i := 0; i < index; i++ {
			cur = cur.Next
		}
		node := &LinkedNode{Val: val}
		node.Next = cur.Next
		cur.Next = node
		list.Size++
	}
}

func Constructor() MyLinkedList {
	return MyLinkedList{
		LinkedNode: &LinkedNode{},
		Size:       0,
	}
}

// Get 获取链表中第 index 个节点的值。如果索引无效，则返回-1。
func (list *MyLinkedList) Get(index int) int {
	if index < 0 || list == nil || index >= list.Size {
		return -1
	}
	cur := list.LinkedNode
	for i := 0; i < index; i++ {
		cur = cur.Next
	}
	return cur.Val
}

// AddAtHead 在链表的第一个元素之前添加一个值为 val 的节点。插入后，新节点将成为链表的第一个节点。
func (list *MyLinkedList) AddAtHead(val int) {
	head := &LinkedNode{Val: val}
	head.Next = list.LinkedNode
	list.LinkedNode = head
	list.Size += 1
}

// AddAtTail 将值为 val 的节点追加到链表的最后一个元素。
func (list *MyLinkedList) AddAtTail(val int) {
	last := &LinkedNode{Val: val}
	for list.LinkedNode.Next != nil {
		if list.LinkedNode.Next.Next == nil {
			list.LinkedNode.Next = last
			break
		}
	}
	list.Size += 1
}

// AddAtIndex 在链表中的第 index 个节点之前添加值为 val  的节点。如果 index 等于链表的长度，则该节点将附加到链表的末尾。如果 index 大于链表长度，则不会插入节点。

// DeleteAtIndex 如果索引 index 有效，则删除链表中的第 index 个节点。
func (list *MyLinkedList) DeleteAtIndex(index int) {
	if index < list.Size {
		cur := list
		for i := 0; i < index; i++ {
			cur.LinkedNode = cur.LinkedNode.Next
		}
		if cur.LinkedNode.Next != nil {
			cur.LinkedNode = cur.LinkedNode.Next
		}
		list.Size -= 1
	}
}
