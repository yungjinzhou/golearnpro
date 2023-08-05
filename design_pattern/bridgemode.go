package main

import "fmt"

// Shape 绘制形状的抽象接口
type Shape interface {
	Draw() string
}

// Color 涂抹颜色的实现部分的接口
type Color interface {
	Paint() string
}

// Red 实现具体的实现部分
type Red struct {
}

func (r *Red) Paint() string {
	return "红色"
}

// Circle  实现具体的抽象部分
type Circle struct {
	color Color
}

func (c *Circle) Draw() string {
	if c.color == nil {
		return ""
	}
	return "Circle: " + c.color.Paint()
}

func (c *Circle) SetColor(color Color) {
	c.color = color
}

func main() {
	red := &Red{}
	circle := &Circle{}

	circle.SetColor(red)
	res := circle.Draw()
	fmt.Println(res)
}
