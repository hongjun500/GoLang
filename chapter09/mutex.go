// @author hongjun500
// @date 2023/8/4 15:19
// @tool ThinkPadX1隐士
// Created with 2022.2.Goland
// Description: sync.Mutex 实现并发安全的计数器

package main

import "sync"

// SafeCounter 并发安全的计数器
type SafeCounter struct {
	mux sync.Mutex // 保护下面的 map: 零值即可用，不要拷贝已用过的 mutex
	v   map[string]int
}

// NewSafeCounter 创建一个新的 SafeCounter，用 make 初始化 map,否则写入时 map 是 nil 会 panic
func NewSafeCounter() *SafeCounter {
	return &SafeCounter{v: make(map[string]int)}
}

// Inc 增加给定 key 的计数器的值
func (c *SafeCounter) Inc(key string) {
	// Lock 之后同一时刻只有一个 goroutine 能访问 c.v
	c.mux.Lock()
	defer c.mux.Unlock()
	c.v[key]++
}

// Value 返回给定 key 的计数器的当前值
func (c *SafeCounter) Value(key string) int {
	// Lock 之后同一时刻只有一个 goroutine 能访问 c.v
	c.mux.Lock()
	// 解锁
	defer c.mux.Unlock()
	return c.v[key]
}
