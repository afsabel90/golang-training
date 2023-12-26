package main

import "fmt"

type cost struct {
	day   int
	value float64
}

func getCostsBydDay(costs []cost) []float64 {

	costsByDay := []float64{}
	for i := 0; i < len(costs); i++ {
		day := costs[i].day
		cost := costs[i].value
		for len(costsByDay) <= day {
			costsByDay = append(costsByDay, 0.0)
		}
		costsByDay[day] += cost
	}

	return costsByDay
}

func main() {
	costs := []cost{
		{0, 4.0},
		{1, 2.1},
		{1, 3.1},
		{5, 2.5},
	}

	fmt.Println(getCostsBydDay(costs))
}
