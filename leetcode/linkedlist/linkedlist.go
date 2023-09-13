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

// 题意：删除链表中等于给定值 val 的所有节点。
// removeElements 移除链表元素
func removeElements(head *LinkedNode, val int) *LinkedNode {
	if head == nil {
		return head
	}
	// 移除头部元素
	for head != nil && head.Val == val {
		head = head.Next
	}
	cur := head
	// 移除中间元素
	for cur != nil && cur.Next != nil {
		if cur.Next.Val == val {
			cur.Next = cur.Next.Next
		} else {
			cur = cur.Next
		}
	}
	return head
}

// MyLinkedList 单向链表
type MyLinkedList struct {
	LinkedNode *LinkedNode
	Size       int
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
	for i := 0; i <= index; i++ {
		cur = cur.Next
	}
	return cur.Val
}

// AddAtHead 在链表的第一个元素之前添加一个值为 val 的节点。插入后，新节点将成为链表的第一个节点。
func (list *MyLinkedList) AddAtHead(val int) {
	/*head := &LinkedNode{Val: val}
	head.Next = list.LinkedNode
	list.LinkedNode = head
	list.Size += 1*/
	list.AddAtIndex(0, val)
}

// AddAtTail 将值为 val 的节点追加到链表的最后一个元素。
func (list *MyLinkedList) AddAtTail(val int) {
	/*last := &LinkedNode{Val: val}
	for list.LinkedNode.Next != nil {
		if list.LinkedNode.Next.Next == nil {
			list.LinkedNode.Next = last
			break
		}
	}
	list.Size += 1*/
	last := &LinkedNode{Val: val}
	list.AddAtIndex(list.Size, last.Val)
}

// AddAtIndex 在链表中的第 index 个节点之前添加值为 val  的节点。如果 index 等于链表的长度，则该节点将附加到链表的末尾。如果 index 大于链表长度，则不会插入节点。
func (list *MyLinkedList) AddAtIndex(index int, val int) {
	/*if index > list.Size {
		return
	}
	if index < 0 {
		index = 0
	}
	cur := list.LinkedNode
	for i := 0; i < index; i++ {
		cur = cur.Next
	}
	node := &LinkedNode{Val: val}
	// node.Next = cur.Next
	cur.Next = node
	list.Size++*/

	if index > list.Size {
		return
	}
	if index < 0 {
		index = 0
	}
	list.Size++
	pred := list.LinkedNode
	for i := 0; i < index; i++ {
		pred = pred.Next
	}
	toAdd := &LinkedNode{val, pred.Next}
	// 由于 pred 是 list.LinkedNode 的引用，所以 pred.Next = toAdd 会影响到 list.LinkedNode.Next
	pred.Next = toAdd
}

// DeleteAtIndex 如果索引 index 有效，则删除链表中的第 index 个节点。
func (list *MyLinkedList) DeleteAtIndex(index int) {
	if index < list.Size {
		cur := list.LinkedNode
		for i := 0; i < index; i++ {
			// 找到要删除节点的前一个节点
			cur = cur.Next
		}
		println(cur.Val)
		/*if cur.Next != nil {
			// 将当前节点的下一个节点指向下下个节点，即为删除当前节点的下一个节点
			cur.Next = cur.Next.Next
		}
		list.Size -= 1*/
	}
}

// 题意：反转一个单链表。
// reverseList 反转链表 双指针
func reverseList(head *LinkedNode) *LinkedNode {
	if head == nil {
		return head
	}
	cur := head
	var pre *LinkedNode
	for cur != nil {
		// 保存当前节点的下一个节点
		next := cur.Next
		// 将当前节点的下一个节点指向前一个节点
		cur.Next = pre
		pre = cur
		// 此时当前节点已经为下一个，一个一个遍历
		cur = next
	}
	return pre
}

// reverseList2 反转链表 递归
func reverseList2(head *LinkedNode) *LinkedNode {
	if head == nil || head.Next == nil {
		return head
	}
	p := reverseList2(head.Next)
	head.Next.Next = head
	head.Next = nil
	return p
}
