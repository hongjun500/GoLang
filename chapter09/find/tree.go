// @author hongjun500
// @date 2023/8/4 14:11
// @tool ThinkPadX1隐士
// Created with 2022.2.Goland
// Description: 等价二叉树（channel 遍历比较）

package find

import (
	"fmt"
	"math/rand"
)

type Tree struct {
	Left  *Tree
	Value int
	Right *Tree
}

// Walk 中序遍历，把所有节点值发送到 ch，遍历完成后关闭 ch。
// close 是关键：让接收方能用 comma-ok 感知"发完了"，从而避免硬编码次数
func Walk(t *Tree, ch chan int) {
	var walk func(*Tree)
	walk = func(n *Tree) {
		if n == nil {
			return
		}
		walk(n.Left)
		ch <- n.Value
		walk(n.Right)
	}
	walk(t)
	close(ch)
}

// Same 判断两棵树的中序序列是否完全一致。
// 注意：这里给通道设了足够大的缓冲，确保即使提前 return，
// Walk 也不会因发送阻塞而泄漏 goroutine（生产环境更推荐用 context 取消）
func Same(t1, t2 *Tree) bool {
	ch1, ch2 := make(chan int, 16), make(chan int, 16)
	go Walk(t1, ch1)
	go Walk(t2, ch2)
	
	for {
		v1, ok1 := <-ch1
		v2, ok2 := <-ch2
		// 一方结束/另一方没结束，或值不同 -> 不相同
		if ok1 != ok2 || v1 != v2 {
			return false
		}
		// 两边同时结束 -> 完全相同
		if !ok1 {
			return true
		}
	}
}

func New(k int) *Tree {
	var t *Tree
	for _, v := range rand.Perm(10) {
		t = insert(t, (1+v)*k)
	}
	return t
}

func insert(t *Tree, v int) *Tree {
	if t == nil {
		return &Tree{nil, v, nil}
	}
	if v < t.Value {
		t.Left = insert(t.Left, v)
	} else {
		t.Right = insert(t.Right, v)
	}
	return t
}

func (t *Tree) String() string {
	if t == nil {
		return "()"
	}
	s := ""
	if t.Left != nil {
		s += t.Left.String() + " "
	}
	s += fmt.Sprint(t.Value)
	if t.Right != nil {
		s += " " + t.Right.String()
	}
	return "(" + s + ")"
}
