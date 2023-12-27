package main

import "fmt"

func add(x, y int) int {
	return x + y
}

func customOperation(x, y int, operation func(int, int) int) int {
	return operation(x, y)
}

func getAddOperation() func(int, int) int {
	return add
}

func doubleOperation(operation func(int, int) int) func(int, int) int {
	return func(x, y int) int {
		return operation(x, y) + operation(x, y)
	}
}

func main() {
	result := customOperation(5, 5, getAddOperation())

	fmt.Println("Custom operation result: ", result)

	result = doubleOperation(add)(5, 5)

	fmt.Println("Double Custom operation result: ", result)
}
