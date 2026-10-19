package main

import "fmt"

func alternatingSum(nums []int) int {

	difference := 0

	for i := 0; i < len(nums); i++{
		if i%2 == 0 {
			difference += nums[i]
		} else {
			difference -= nums[i]
		}
	}

	return difference

}

func main() {
	fmt.Println(alternatingSum([]int{1, 3, 5, 7}))
}
