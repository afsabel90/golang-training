package main

import "fmt"

func getPrimes() [6]int {
	return [6]int{2, 3, 5, 7, 11, 13}
}

func getPrimesSlice() []int {
	return []int{1, 2, 3, 5, 7, 11, 13}
}

func sum(num ...int) int {
	total := 0

	for i := 0; i < len(num); i++ {
		total += num[i]
	}

	return total
}

func main() {

	fmt.Println(getPrimes())
	fmt.Println(getPrimesSlice())
	fmt.Println(getPrimesSlice()[1:])
	fmt.Println(sum(1, 2, 3, 4, 5))
}
