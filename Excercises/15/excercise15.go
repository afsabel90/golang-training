package main

import (
	"fmt"
)

type craftie struct {
	name string
	role string
}

func work(c craftie, triggerError bool) error {
	if triggerError {
		return fmt.Errorf("craftie %v had a problem", c.name)
	} else {
		fmt.Printf("Craftie %v is working!", c.name)
	}
	return nil
}

// Entry point

func main() {
	e := work(craftie{
		name: "Agustin",
		role: "Developer",
	}, false)

	if e != nil {
		fmt.Println(e.Error())
	}

	e = work(craftie{
		name: "Agustin",
		role: "Developer",
	}, true)

	if e != nil {
		fmt.Println(e.Error())
	}
}
