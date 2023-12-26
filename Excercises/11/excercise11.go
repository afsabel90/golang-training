package main

import "fmt"

type rect struct {
	width  int
	height int
}

func (r rect) area() int {
	return r.width * r.height
}

// Entry point

func main() {

	square := rect{
		width:  5,
		height: 5,
	}

	fmt.Println(square.area())
}
