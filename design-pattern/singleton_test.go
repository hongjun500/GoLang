// @author hongjun500
// @date 2023/9/14 14:59
// @tool ThinkPadX1隐士
// Created with 2022.2.Goland
// Description:

package main

import (
	"testing"

	"github.com/go-playground/assert/v2"
)

func TestNewSingleton(t *testing.T) {
	assert.Equal(t, NewSingleton(), NewSingleton())
}

func BenchmarkNewSingleton(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			_ = NewSingleton()
		}
	})
	// for i := 0; i < b.N; i++ {
	// 	_ = NewSingleton()
	// }
}

func TestNewLazySingleton(t *testing.T) {
	assert.Equal(t, NewLazySingleton(), NewLazySingleton())
}

func BenchmarkNewLazySingleton(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			_ = NewLazySingleton()
		}
	})
	// for i := 0; i < b.N; i++ {
	// 	_ = NewLazySingleton()
	// }
}
