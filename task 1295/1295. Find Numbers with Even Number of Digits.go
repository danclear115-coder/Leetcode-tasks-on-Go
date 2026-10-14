package main

import "fmt"

func getNumsCounter(n int) (counter int) {

	for n > 0 {
		counter++
		n = n / 10
	}

	return

}

func findNumbers(nums []int) int {

	counter := 0

	for _, num := range nums {
		if getNumsCounter(num)%2 == 0 {
			counter++
		}
	}

	return counter

}

func main() {
	fmt.Println(findNumbers([]int{12,345,2,6,7896}))
}
