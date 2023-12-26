package main

import "fmt"

// Entry point

func main() {

	var username string = "craftie"
	var password string = "1234567890"
	fmt.Println("Authorization: Basic", username+":"+password)
}
