package main

import "fmt"

func sortSlice(slice []int) (sortedSlice []int) {

	for len(slice) > 0 {

		minElement, minIndex := slice[0], 0

		for i := 0; i < len(slice); i++ {
			if slice[i] < minElement {
				minElement, minIndex = slice[i], i
			}
		}

		sortedSlice = append(sortedSlice, minElement)
		slice = append(slice[:minIndex], slice[minIndex+1:]...)

	}

	return

}

func searchInsert(nums []int, target int) int {

	if target < nums[0] {
		return 0
	}

	position := 0

	sortedNums := sortSlice(nums)
	fmt.Println(sortedNums)

	for i := 0; i < len(sortedNums); i++ {
		if target <= sortedNums[i] {
			return i
		} else if i+1 < len(sortedNums) && sortedNums[i]+1 != sortedNums[i+1] && sortedNums[i+1] >= target {
			return i + 1
		} else if i+1 == len(sortedNums) && sortedNums[i] < target {
			return i + 1
		}
	}

	return position

}

func main() {
	fmt.Println(searchInsert([]int{1, 3, 5, 6}, 5))
}
