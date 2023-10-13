// @author hongjun500
// @date 2023/10/13 14:08
// @tool ThinkPadX1隐士
// Created with 2022.2.Goland
// Description:

package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestColorSquare_Draw(t *testing.T) {
	square := Square{}
	colorSquare := NewColorSquare(square, "red")
	draw := colorSquare.Draw()
	assert.Equal(t, "square, color is red", draw)
}
