package main

import "fmt"

func isUnique(nums []int, num int, numIndex int) bool {

	for i := 0; i < len(nums); i++ {
		if nums[i] == num && i != numIndex {
			return false
		}
	}

	return true

}

func sumOfUnique(nums []int) int {
    
	sum := 0

	for i := 0; i < len(nums); i++ {
		if isUnique(nums, nums[i], i) {
			sum += nums[i]
		}
	}

	return sum

}

func main() {
	fmt.Println(sumOfUnique([]int{1, 2, 3, 2}))
}
