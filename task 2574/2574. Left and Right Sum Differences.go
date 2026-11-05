package main

import "fmt"

func leftRightDifference(nums []int) []int {

	result := []int{}

	for a := 0; a < len(nums); a++ {

		leftSum, rightSum := 0, 0

		for b := 0; b < a; b++ {
			leftSum += nums[b]
		}

		for c := a + 1; c < len(nums); c++ {
			rightSum += nums[c]
		}

		if leftSum - rightSum < 0 {
			result = append(result, (leftSum - rightSum) * -1)
		} else {
			result = append(result, leftSum - rightSum)
		}

	}

	return result

}

func main() {
	fmt.Println(leftRightDifference([]int{10, 4, 8, 3}))
}
