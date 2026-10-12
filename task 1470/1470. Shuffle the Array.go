package main

import "fmt"

func shuffle(nums []int, n int) []int {

	nums1, nums2, result, counter := []int{}, []int{}, []int{}, 0

	for counter < n {
		nums1 = append(nums1, nums[counter])
		counter++
	}

	for counter < len(nums) {
		nums2 = append(nums2, nums[counter])
		counter++
	}

	for i := 0; i < n; i++ {
		result = append(result, nums1[i])
		result = append(result, nums2[i])
	}

	return result

}

func main() {
	fmt.Println(shuffle([]int{2, 5, 1, 3, 4, 7}, 3))
}
