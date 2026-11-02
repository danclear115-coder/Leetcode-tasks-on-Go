package main

import "fmt"

func sumOfSquares(nums []int) int {

	sum, newNum := 0, []int{0}
	newNum = append(newNum, nums...)

	for i := 1; i < len(newNum); i++ {
		if (len(newNum) - 1) % i == 0 {
			sum += newNum[i] * newNum[i]
		}
	}

	return sum

}

func main() {
	fmt.Println(sumOfSquares([]int{1, 2, 3, 4}))
}
