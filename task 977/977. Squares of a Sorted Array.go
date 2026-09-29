package main

import "fmt"

func sortNums(nums []int) (sortedNums []int) {

	for len(nums) > 0 {

		min, indexOfMin := nums[0], 0

		for i := 0; i < len(nums); i++ {
			if nums[i] < min {
				min, indexOfMin = nums[i], i
			}
		}

		sortedNums = append(sortedNums, min)
		nums = append(nums[:indexOfMin], nums[indexOfMin+1:]...)

	}

	return

}

func sortedSquares(nums []int) []int {

	squareNums := []int{}

	for _, num := range nums {
		squareNums = append(squareNums, num*num)
	}

	sortedSquareNums := sortNums(squareNums)

	return sortedSquareNums

}

func main() {
	fmt.Println(sortedSquares([]int{-4, -1, 0, 3, 10}))
}
