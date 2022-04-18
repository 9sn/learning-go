package main

import (
	"fmt"
	"math"
	
)

type Shape interface {
	area() float64
}

type circle struct {
	x, y, radius float64
}

type rectangle struct {
	width, height float64
}

func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func getArea(s Shape) float64 {
	return s.area()
}

func (r rectangle) area() float64 {
	return r.width * r.height
}

func main() {
	circle := circle{x:0,y:0,radius:5}
	rectangle := rectangle{width: 10, height: 5}
	
	fmt.Printf("Circle Area: %f\n", getArea(circle))
	fmt.Printf("Rectangle Area: %f\n", getArea(rectangle))
}