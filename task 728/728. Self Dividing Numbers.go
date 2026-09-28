package main

import "fmt"

func isSelfDiv(n int) bool {

	number, nums := n, []int{}

	for number > 0 {
		nums = append(nums, number % 10)
		number = number / 10
	}

	for _, div := range nums {
		if div == 0 || n % div != 0 {
			return false
		}
	}

	return true

}

func selfDividingNumbers(left int, right int) []int {
    
	selfDivs := []int{}

	for i := left; i <= right; i++ {
		if isSelfDiv(i) {
			selfDivs = append(selfDivs, i)
		}
	}

	return selfDivs

}

func main() {
	fmt.Println(isSelfDiv(12))
	fmt.Println(selfDividingNumbers(1, 22))
}
