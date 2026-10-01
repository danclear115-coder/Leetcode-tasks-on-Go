package main

import "fmt"

func isInclude(nums []int, num int) bool {

	for _, element := range nums {
		if num == element {
			return true
		}
	}

	return false

}

func getMinAndMax(nums []int) (int, int) {

	max, min := nums[0], nums[0]

	for _, num := range nums {
		if num < min {
			min = num
		}
	}

	for _, num := range nums {
		if num > max {
			max = num
		}
	}

	return min, max

}

func getFullNums(nums []int) (fullSortedNums []int) {

	min, max := getMinAndMax(nums)
	fullNums := []int{min}

	fmt.Println(nums, min, max)

	for i := min + 1; i <= max; i++ {
		fullNums = append(fullNums, i)
	}

	return fullNums

}

func findMissingElements(nums []int) []int {

	result, fullNums := []int{}, getFullNums(nums)

	for _, num := range fullNums {
		if !isInclude(nums, num) {
			result = append(result, num)
		}
	}

	return result

}

func main() {
	fmt.Println(findMissingElements([]int{1, 4, 2, 5}))
}
