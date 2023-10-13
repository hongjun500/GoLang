// @author hongjun500
// @date 2023/10/13 14:04
// @tool ThinkPadX1隐士
// Created with 2022.2.Goland
// Description: 装饰者模式

package main

type IDraw interface {
	Draw() string
}

// Square 正方形
type Square struct {
}

func (s Square) Draw() string {
	return "square"
}

type ColorSquare struct {
	square IDraw
	color  string
}

func NewColorSquare(square IDraw, color string) ColorSquare {
	return ColorSquare{square: square, color: color}
}

func (c ColorSquare) Draw() string {
	return c.square.Draw() + ", color is " + c.color
}
