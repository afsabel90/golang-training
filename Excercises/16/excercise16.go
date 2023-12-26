package main

import "fmt"

func bulkSend(numMessages int) float64 {
	cost := 0.0
	for i := 0; i < numMessages; i++ {
		cost += 1.00 + (0.01 * float64(i))
	}

	return cost
}

// Entry point

func main() {
	fmt.Println("Total cost: ", bulkSend(5))
}
