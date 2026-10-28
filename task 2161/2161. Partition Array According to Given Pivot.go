package main

import "fmt"

func pivotArray(nums []int, pivot int) []int {

	midNums, heighNums, result := []int{}, []int{}, []int{}

	for a := 0; a < len(nums); a++ {
		if nums[a] < pivot {
			result = append(result, nums[a])
		} else if nums[a] == pivot {
			midNums = append(midNums, nums[a])
		} else {
			heighNums = append(heighNums, nums[a])
		}
	}

	result = append(result, midNums...)
	result = append(result, heighNums...)

	return result

}

func main() {
	fmt.Println(pivotArray([]int{1, 4, 5, 12, 34, 10, 19, 15, 16, 4}, 10))
}
