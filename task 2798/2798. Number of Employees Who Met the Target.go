package main

import "fmt"

func numberOfEmployeesWhoMetTarget(hours []int, target int) int {
	
	counter := 0

	for _, timeMark := range hours {
		if timeMark >= target {
			counter++
		}
	}

	return counter

}

func main() {
	fmt.Println(numberOfEmployeesWhoMetTarget([]int{0, 1, 2, 3, 4}, 2))
}
