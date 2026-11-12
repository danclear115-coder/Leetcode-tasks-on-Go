package main

import (
	"fmt"
	"slices"
)

func isPairInclude(nums []int) bool {

	for a := 0; a < len(nums); a++ {
		for b := a + 1; b < len(nums); b++ {
			if nums[a] == nums[b] {
				return true
			}
		}
	}

	return false

}

func getSliceAfterpairDelete(nums []int) (newNums []int) {

	for a := 0; a < len(nums); a++ {
		for b := a + 1; b < len(nums); b++ {
			if nums[a] == nums[b] {
				nums = slices.Delete(nums, b, b+1)
				nums = slices.Delete(nums, a, a+1)
				newNums = nums
				return
			}
		}
	}

	newNums = nums
	return

}

func numberOfPairs(nums []int) []int {

	pairCounter := 0

	for isPairInclude(nums) {
		nums = getSliceAfterpairDelete(nums)
		pairCounter++
	}

	return []int{pairCounter, len(nums)}

}

func main() {
	fmt.Println(numberOfPairs([]int{1, 1}))
}
