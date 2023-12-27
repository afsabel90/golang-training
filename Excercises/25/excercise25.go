package main

import "time"

func main() {
	ch := make(chan int)

	go func(ch chan int) {
		time.Sleep(time.Millisecond * 1000)
		ch <- 10
	}(ch)

	v := <-ch

	println("Int Channel value: ", v)
}
