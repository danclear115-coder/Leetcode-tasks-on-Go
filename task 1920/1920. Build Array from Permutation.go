package main

import "fmt"

func buildArray(nums []int) []int {
 
	result := []int{}

	for i := 0; i < len(nums); i++ {
		result = append(result, nums[nums[i]])
	}

	return result

}

func main() {
	fmt.Println(buildArray([]int{0, 2, 1, 5, 3, 4}))
}