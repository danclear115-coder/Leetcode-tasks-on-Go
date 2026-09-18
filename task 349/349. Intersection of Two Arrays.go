package main

import "fmt"

func isInclude(nums []int, element int) bool {

	for _, num := range nums {
		if num == element {
			return true
		}
	}

	return false

}

func intersection(nums1 []int, nums2 []int) []int {

	result := []int{}

	for _, num1 := range nums1 {
		for _, num2 := range nums2 {
			if num1 == num2 && !isInclude(result, num1) {
				result = append(result, num1)
			}
		}
	}

	return result

}

func main() {
	fmt.Println(intersection([]int{4, 9, 5}, []int{}))
}
