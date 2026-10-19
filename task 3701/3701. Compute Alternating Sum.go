package main

import "fmt"

func alternatingSum(nums []int) int {

	sumOfEven, sumOfNotEven := 0, 0

	for i := 0; i < len(nums); i++{
		if i%2 == 0 {
			sumOfEven += nums[i]
		} else {
			sumOfNotEven += nums[i]
		}
	}

	return sumOfEven - sumOfNotEven

}

func main() {
	fmt.Println(alternatingSum([]int{1, 3, 5, 7}))
}
