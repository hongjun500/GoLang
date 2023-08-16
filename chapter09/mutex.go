// @author hongjun500
// @date 2023/8/4 15:19
// @tool ThinkPadX1隐士
// Created with 2022.2.Goland
// Description: 互斥锁

package main

import "sync"

type SafeCounter struct {
	v   map[string]int
	mux sync.Mutex
}

// Inc 增加给定 key 的计数器的值
func (c *SafeCounter) Inc(key string) {
	// Lock 之后同一时刻只有一个 goroutine 能访问 c.v
	c.mux.Lock()
	c.v[key]++
	c.mux.Unlock()
}

// Value 返回给定 key 的计数器的当前值
func (c *SafeCounter) Value(key string) int {
	// Lock 之后同一时刻只有一个 goroutine 能访问 c.v
	c.mux.Lock()
	// 解锁
	defer c.mux.Unlock()
	return c.v[key]
}
