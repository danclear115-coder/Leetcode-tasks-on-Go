package main

import "fmt"

func getSum(num int) (sum int) {

	for num > 0 {
		sum += num % 10
		num = num / 10
	}

	return sum

}

func getMin(nums []int) int {

	min := nums[0]

	for _, num := range nums {
		if num < min {
			min = num
		}
	}

	return min

}

func minElement(nums []int) int {

	sums := []int{}

	for _, num := range nums {
		sums = append(sums, getSum(num))
	}

	return getMin(sums)

}

func main() {
	fmt.Println(minElement([]int{999,19,199}))
}
