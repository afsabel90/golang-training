package main

import "fmt"

// Entry point

func main() {

	len := 10

	if len >= 5 {
		fmt.Println("Lenght is enough")
	} else if len <= 1 {
		fmt.Println("Lenght is terribly short")
	} else {
		fmt.Println("Lenght is not enough")
	}
}
