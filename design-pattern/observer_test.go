// @author hongjun500
// @date 2023/10/16 15:43
// @tool ThinkPadX1隐士
// Created with 2022.2.Goland
// Description:

package main

import "testing"

func TestObserver(t *testing.T) {
	var sub ISubject = &Subject{}
	// sub := &Subject{}
	sub.Register(&Observer1{})
	sub.Register(&Observer2{})
	sub.Notify("test")
}
