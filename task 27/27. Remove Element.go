package main

import "fmt"

func removeElement(nums []int, val int) int {

	index := 0

	for a := 0; a < len(nums); a++ {
		if nums[a] != val {
			nums[index] = nums[a]
			index++
		}
	}

	fmt.Println(nums)

	return index

}

func main() {
	fmt.Println(removeElement([]int{3, 2, 2, 3}, 3))
}
