package main

import "fmt"

func test(a *int) int {
	return *a * *a
}

func main() {
	x := 50
	fmt.Println(test(&x))
}
