package main

import "fmt"

func isTwiceElement(nums []int, num int) bool {

	counter := 0

	for _, element := range nums {
		if element == num {
			counter++
		}
	}

	if counter >= 2 {
		return true
	}

	return false

}

func isInclude(nums []int, num int) bool {

	for _, element := range nums {
		if element == num {
			return true
		}
	}

	return false

}

func getSneakyNumbers(nums []int) []int {

	result := []int{}

	for _, num := range nums {
		if isTwiceElement(nums, num) && !isInclude(result, num) {
			result = append(result, num)
		}
	}

	return result

}

func main() {
	fmt.Println(getSneakyNumbers([]int{0,1,1,0}))
}
