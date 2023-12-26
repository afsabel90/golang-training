package main

import "fmt"

type craftie struct {
	Name string
	Role string
	Age  int
}

// Entry point

func main() {

	agus := craftie{
		Name: "Agustin",
		Role: "Developer",
		Age:  33,
	}

	fmt.Println(agus)

	agus2 := struct {
		Name string
		Role string
		Age  int
	}{
		Name: "Agustin",
		Role: "Developer",
		Age:  33,
	}

	fmt.Println(agus2)
}
