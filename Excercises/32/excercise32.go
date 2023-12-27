package main

import "fmt"

type craftie struct {
	name string
}

type person struct {
	name string
}

type entity interface {
	getName() string
}

func (c craftie) getName() string {
	return c.name
}

func (p person) getName() string {
	return p.name
}

func getEntityName[T entity](e T) string {
	return e.getName()
}

func main() {
	agus := craftie{
		name: "Agus",
	}

	v := getEntityName[craftie](agus)

	fmt.Println(v)

	john := person{
		name: "John",
	}

	v = getEntityName[person](john)

	fmt.Println(v)
}
