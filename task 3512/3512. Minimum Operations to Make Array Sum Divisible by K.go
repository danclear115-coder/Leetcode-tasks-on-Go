package main

import "fmt"

func minOperations(nums []int, k int) int {
    
	sum, count := 0, 0

	for i := 0; i < len(nums); i++ {
		sum += nums[i]
	}

	for {
		if sum % k == 0 {
			return count
		} else {
			sum, count = sum - 1, count + 1
		}
	}

}

func main() {
	fmt.Println(minOperations([]int{3, 9, 7}, 5))
}
