package main

import "fmt"

func deferMe() {
	fmt.Println("Show at the end")
}

func main() {
	defer deferMe()

	fmt.Println("I'll go first!")
}
