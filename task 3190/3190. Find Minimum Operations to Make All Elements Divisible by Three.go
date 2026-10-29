package main

import "fmt"

func getOptimalOperationPath(num int) (optimalOperation int) {

	minusOps, plusOps, num1, num2 := 0, 0, num, num

	for ;num1 % 3 != 0; num1-- {
		minusOps++
	}

	for ;num2 % 3 != 0; num2++ {
		plusOps++
	}

	if minusOps < plusOps {
		return minusOps
	} else if minusOps > plusOps {
		return plusOps
	}

	return minusOps

}

func minimumOperations(nums []int) int {
    
	counter := 0

	for _, number := range nums {
		counter += getOptimalOperationPath(number)
	}

	return counter

}

func main() {
	fmt.Println(minimumOperations([]int{3, 6, 9}))
}
