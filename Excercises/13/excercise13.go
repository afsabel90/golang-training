package main

import "fmt"

type shape interface {
	area() float64
}

type rect struct {
	width  float64
	height float64
}

func (r rect) area() float64 {
	return r.width * r.height
}

// Entry point

func main() {
	var s shape = rect{
		width:  20,
		height: 20,
	}

	r, ok := s.(rect)

	fmt.Println(r, ok)
}
