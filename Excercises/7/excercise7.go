package main

import "fmt"

// Entry point

func main() {

	x := 5
	y := &x

	x = 8

	fmt.Println("X:", x, "Y:", *y)
}
