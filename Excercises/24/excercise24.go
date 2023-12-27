package main

import (
	"fmt"
	"time"
)

func awaitMe() {
	time.Sleep(time.Microsecond * 250)
	fmt.Println("I'll be printed at the end")
}

func main() {
	go awaitMe()
	fmt.Println("I'll go first!")
	time.Sleep(time.Microsecond * 250)
}
