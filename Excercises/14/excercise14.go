package main

import "fmt"

type craftieError struct {
	name string
}

type craftie struct {
	name string
	role string
}

func (e craftieError) Error() string {
	return fmt.Sprintf("Craftie %v had a problem!", e.name)
}

func work(c craftie, triggerError bool) error {
	if triggerError {
		return craftieError{name: c.name}
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
