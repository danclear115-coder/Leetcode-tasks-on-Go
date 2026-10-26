package main

import "fmt"

func concatWithReverse(nums []int) []int {
    
	for i := len(nums) - 1; i >= 0; i-- {
		nums = append(nums, nums[i])
	}

	return nums

}

func main() {
	fmt.Println(concatWithReverse([]int{1, 2, 3}))
}
