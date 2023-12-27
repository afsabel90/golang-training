package main

import "fmt"

type craftie struct {
	Name string
	Role string
	Age  int
}

func main() {

	ages := make(map[string]int)

	ages["Agus"] = 33

	ages2 := make(map[int]int)

	ages2[0] = 33

	crafties := make(map[string]craftie)

	crafties["Agus"] = craftie{
		Name: "Agus",
		Role: "Developer",
		Age:  33,
	}

	fmt.Println(crafties["Agus"])

	randomName := "Fer"

	elem, ok := crafties[randomName]

	if !ok {
		fmt.Println("Craftie ", randomName, " does not exist")
	} else {
		fmt.Println(elem)
	}
}
