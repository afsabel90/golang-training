package main

import (
	"fmt"
	"math/rand"
)

func main() {

	chInts := make(chan int)
	chStrings := make(chan string)

	go func() {
		for {
			number := rand.Int()

			if number%2 == 0 {
				chInts <- number
			} else {
				chStrings <- "Odd found"
				close(chInts)
				close(chStrings)
				break
			}
		}
	}()

	ok := true
	var i int
	var s string

	for ok {
		select {
		case i, ok = <-chInts:
			if ok {
				fmt.Println(i)
			}
		case s, ok = <-chStrings:
			if ok {
				fmt.Println(s)
			}
		}
	}
}
