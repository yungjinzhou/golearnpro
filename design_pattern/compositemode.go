package main

import (
	"fmt"
)

type Graphic interface {
	Draw()
}

type Point struct {
	x, y int
}

func (p *Point) String() string {
	return fmt.Sprintf("点(%d, %d)", p.x, p.y)
}

func (p *Point) Draw() {
	fmt.Println(p.String())
}

type Line struct {
	p1, p2 *Point
}

func (l *Line) String() string {
	return fmt.Sprintf("线段[%s, %s]", l.p1, l.p2)
}

func (l *Line) Draw() {
	fmt.Println(l.String())

}

type Picture struct {
	children []Graphic
}

func (pi *Picture) Draw() {
	fmt.Println("----复合图形-----")
	for _, g := range pi.children {
		g.Draw()
		fmt.Println("----复合图形-----")
	}
}

func (pi *Picture) add(g Graphic) {
	pi.children = append(pi.children, g)
}

func main() {
	p1 := &Point{2, 3}
	l1 := &Line{&Point{3, 4}, &Point{6, 7}}
	l2 := &Line{&Point{1, 5}, &Point{2, 8}}
	pic1 := &Picture{}
	pic1.add(p1)
	pic1.add(l1)
	pic1.add(l2)

	p2 := &Point{4, 4}
	l3 := &Line{&Point{1, 1}, &Point{0, 0}}
	pic2 := &Picture{}
	pic2.add(p2)
	pic2.add(l3)

	pic := &Picture{}
	pic.add(pic1)
	pic.add(pic2)

	pic.Draw()
}
