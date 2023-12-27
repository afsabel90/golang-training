package main

import "fmt"

func main() {

	len := 5
	ch := make(chan int)

	go func() {
		for i := 0; i < len; i++ {
			ch <- i
		}
		close(ch)
	}()

	for item := range ch {
		fmt.Println("Item #", item, " received!")
	}
}
