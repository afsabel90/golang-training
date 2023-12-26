package main

import (
	"errors"
	"fmt"
)

func divide(dividend, divisor int) (int, error) {
	if divisor == 0 {
		return 0, errors.New("can't devide by zero")
	}
	return dividend / divisor, nil
}

// Entry point

func main() {

	var result, _ = divide(5, 0)
	fmt.Println("Result : ", result)
}
