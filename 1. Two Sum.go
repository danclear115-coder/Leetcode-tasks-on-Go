package main

import "fmt"

func twoSum(nums []int, target int) []int {

	result := []int{}

	for a := 0; a < len(nums); a++ {
		for b := a + 1; b < len(nums); b++ {
			if nums[a] + nums[b] == target {
				return []int{a, b}
			}
		}
	}

	return result

}

func main() {

	fmt.Println(twoSum([]int{3, 2, 3}, 6))

}
