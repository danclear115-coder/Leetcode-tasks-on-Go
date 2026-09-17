package main

import "fmt"

func findMin(nums []int) int {

	minElement := nums[0]

	for _, num := range nums {
		if num < minElement {
			minElement = num
		}
	}

	return minElement

}

func findMax(nums []int) int {

	maxElement := nums[0]

	for _, num := range nums {
		if num > maxElement {
			maxElement = num
		}
	}

	return maxElement

}

func findGCD(nums []int) int {

	min, max, maxDiv := findMin(nums), findMax(nums), 1

	for i := 1; i <= max; i++ {
		if max%i == 0 && min%i == 0 && maxDiv < i {
			maxDiv = i
		}
	}

	return maxDiv

}

func main() {
	fmt.Println(findGCD([]int{7,5,6,8,3}))
}
