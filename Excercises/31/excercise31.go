package main

import "fmt"

func split[T any](s []T) ([]T, []T) {
	mid := len(s) / 2
	return s[:mid], s[mid:]
}

func main() {
	test := []int{1, 2, 3, 5, 6, 7, 8, 9}
	a, b := split(test)
	fmt.Println(a, b)
	test2 := []string{"hello", "craftie", "how", "are", "you", "doing?"}
	c, d := split(test2)
	fmt.Println(c, d)
}
