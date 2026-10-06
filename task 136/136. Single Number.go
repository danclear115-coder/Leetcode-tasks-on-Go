package main

import "fmt"

func isDoubleCheck(nums []int, num int) bool {

	counter := 0

	for _, numElement := range nums {
		if numElement == num {
			counter++
		}
	}

	if counter != 2 {
		return false
	}

	return true

}

func singleNumber(nums []int) int {

	for i := 0; i < len(nums); i++ {
		if !isDoubleCheck(nums, nums[i]) {
			return nums[i]
		}
	}

	return 0

}

func main() {

	fmt.Println(singleNumber([]int{1,2,1,2, 5, 6, 6}))

}
