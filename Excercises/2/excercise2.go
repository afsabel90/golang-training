package main

import "fmt"

// Entry point

func main() {

	messages := []string{
		"Hello Craftie!",
		"How are you doing?",
		"I'm fine too",
		"Let's catch up later!",
	}

	numMessages := float64(len(messages))
	cost := 0.2

	totalCost := cost * numMessages

	fmt.Printf("The craftie spent %.2f on text messages", totalCost)

}
