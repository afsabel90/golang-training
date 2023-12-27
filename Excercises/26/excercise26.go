package main

import (
	"fmt"
	"time"
)

type thing struct {
	Name string
}

func waitForThings(things int, somethingChan chan thing) {
	for i := 0; i < things; i++ {
		<-somethingChan
	}

	fmt.Println("All things were received!")
}

func main() {
	ch := make(chan thing)
	go waitForThings(6, ch)

	ch <- thing{
		Name: "Thing1",
	}

	ch <- thing{
		Name: "Thing2",
	}
	ch <- thing{
		Name: "Thing3",
	}
	ch <- thing{
		Name: "Thing4",
	}
	ch <- thing{
		Name: "Thing5",
	}
	ch <- thing{
		Name: "Thing6",
	}

	time.Sleep(time.Millisecond * 200)
}
