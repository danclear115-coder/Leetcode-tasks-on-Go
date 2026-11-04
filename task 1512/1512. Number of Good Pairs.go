package main

import "fmt"

func getTrueIncludeCounter(nums []int, num int, index int) (counter int) {

	for i := index; i < len(nums); i++ {
		if nums[i] == num && i != index {
			counter++
		}		
	}

	return

}

func numIdenticalPairs(nums []int) int {

	counter := 0

	for i := 0; i < len(nums); i++ {
		numCounter := getTrueIncludeCounter(nums, nums[i], i)
		counter += numCounter
	}

	return counter

}

func main() {
	fmt.Println(numIdenticalPairs([]int{1, 2, 3}))
}
