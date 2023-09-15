// @author hongjun500
// @date 2023/9/14 14:24
// @tool ThinkPadX1隐士
// Created with 2022.2.Goland
// Description: 单例模式

package main

import "sync"

type Singleton struct {
}

var singleton *Singleton

// init 初始化单例 饿汉式
func init() {
	singleton = &Singleton{}
}

// NewSingleton 获取单例
func NewSingleton() *Singleton {
	return singleton
}

var lazySingleton *Singleton
var once sync.Once

// NewLazySingleton 懒汉式
func NewLazySingleton() *Singleton {
	if lazySingleton == nil {
		once.Do(func() {
			lazySingleton = &Singleton{}
		})
	}
	return lazySingleton
}
