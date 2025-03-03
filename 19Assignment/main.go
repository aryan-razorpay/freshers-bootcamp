package main

import (
	"fmt"
	"math"
)

type Shape interface {
	Area() float64
}

type Circle struct {
	radius float64
}

type Rectangle struct {
	width, height float64
}

type Square struct {
	side float64
}

func (c Circle) Area() float64 {
	return math.Pi * c.radius * c.radius
}

func (r Rectangle) Area() float64 {
	return r.width * r.height
}

func (s Square) Area() float64 {
	return s.side * s.side
}

func main() {
	circle := Circle{radius: 5}
	rectangle := Rectangle{width: 4, height: 6}
	square := Square{side: 3}
	shapes := []Shape{circle, rectangle, square}

	for _, shape := range shapes {
		fmt.Printf("Area of %T: %.2f\n", shape, shape.Area())
	}
}
