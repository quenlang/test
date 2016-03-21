package main

import (
	"fmt"
	"math"
)

type Point struct {
	x, y float64
}

type NamePoint struct {
	// 由于包含了Point类型，因此继承了Point的方法
	Point
	name string
}

func main() {
	n := &NamePoint{Point{3, 4}, "Pythagoras"}
	fmt.Println(n.Abs())

}

func (p *Point) Abs() float64 {
	return math.Sqrt(p.x*p.x + p.y*p.y)
}

// 类似于面向对象中的方法的重载
func (n *NamePoint) Abs() float64 {
	return n.Point.Abs() * 100
}
