package main

import (
	"fmt"
	"math"
)

// Rectangle is 矩形
type Rectangle struct {
	x1, x2, y1, y2 float64
}

func distance(x1, x2, y1, y2 float64) float64 {
	return (y1 - x1) * (y2 - x2)
}

func (r *Rectangle) area() float64 {
	return distance(r.x1, r.x2, r.y1, r.y2)
}

// Circle is 圆形
type Circle struct {
	x, y, r float64
}

func (c *Circle) area() float64 {
	return math.Pi * c.r * c.r
}

// Shape is 图形（只要有面积属性，那么都属于图形）
type Shape interface {
	area() float64
}

func totalArea(shapes ...Shape) float64 {
	var area float64
	for _, s := range shapes {
		area += s.area()
	}
	return area
}

func main() {
	c := Circle{1, 1, 1}
	r := Rectangle{0, 0, 3, 3}
	fmt.Println(totalArea(&c, &r))
}
