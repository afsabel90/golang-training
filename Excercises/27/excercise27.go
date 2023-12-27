package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	len := 2
	iterLen := 3
	ch := make(chan int, len)

	go func() {
		for i := 0; i < iterLen; i++ {
			ch <- rand.Int()
		}
		close(ch)
		fmt.Println("This will never be displayed...") // len needs to be equal or greater than len for this to be displayed...
	}()

	time.Sleep(time.Millisecond * 1000)
}
