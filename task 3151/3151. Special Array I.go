package main

import "fmt"

func isPar(n int) bool {

	if n % 2 == 0 {
		return true
	}

	return false

}

func isArraySpecial(nums []int) bool {

	for i := 0; i < len(nums); i++ {
		if i + 1 < len(nums) {
			if isPar(nums[i]) == isPar(nums[i + 1]) {
				return false
			}
		}
	}	

	return true 

}

func main() {
	fmt.Println(isArraySpecial([]int{4, 3, 1, 6}))
}