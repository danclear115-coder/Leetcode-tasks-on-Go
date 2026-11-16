package main

import (
	"fmt"
	"slices"
)

func sortToMax(nums []int) (sortedNums []int) {

	for len(nums) > 0 {
		min, index := nums[0], 0
		for i := 0; i < len(nums); i++ {
			if nums[i] < min {
				min, index = nums[i], i
			}
		}
		sortedNums = append(sortedNums, min)
		nums = slices.Delete(nums, index, index + 1)
	}

	return

}

func addedInteger(nums1 []int, nums2 []int) int {
	
	nums1, nums2 = sortToMax(nums1), sortToMax(nums2)

	return (nums2[0] - nums1[0])

}

func main () {
	fmt.Println(addedInteger([]int{2, 6, 4}, []int{9, 7, 5}))
}

