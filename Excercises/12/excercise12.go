package main

import (
	"fmt"
	"math"
)

type shape interface {
	area() float64
	perimeter() float64
}

type rect struct {
	width  float64
	height float64
}

type circle struct {
	radius float64
}

func (r rect) area() float64 {
	return r.width * r.height
}

func (r rect) perimeter() float64 {
	return r.height*2 + r.width*2
}

func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func (c circle) perimeter() float64 {
	return 2 * math.Pi * c.radius
}

func printArea(shape shape) {
	fmt.Println("Area: ", shape.area())
	fmt.Println("Perimeter", shape.perimeter())
}

// Entry point

func main() {

	square := rect{
		width:  5,
		height: 5,
	}

	circle := circle{
		radius: 10,
	}

	printArea(circle)
	printArea(square)
}
