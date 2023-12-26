package main

import "fmt"

type craftie struct {
	Name string
	Role string
	Age  int
	hobby
}

type hobby struct {
	HobbyName  string
	HobbySince int
}

// Entry point

func main() {

	agus := craftie{
		Name: "Agustin",
		Role: "Developer",
		Age:  33,
		hobby: hobby{
			HobbyName:  "Charcuterie",
			HobbySince: 2015,
		},
	}

	agus.HobbyName = "Charcuterie2"

	fmt.Println(agus)
}
