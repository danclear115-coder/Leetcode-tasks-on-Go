package main

import "fmt"

func countPairs(nums []int, target int) int {

	counter := 0

	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i]+nums[j] < target {
				counter++
			}
		}
	}

	return counter

}

func main() {
	fmt.Println(countPairs([]int{-6, 2, 5, -2, -7, -1, 3}, -2))
}
