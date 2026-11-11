package main

import "fmt"

func transformArray(nums []int) []int {
    
    result := []int{}

	for _, num := range nums {
		if num%2 == 0 {
			result = append([]int{0}, result...)
		} else {
			result = append(result, 1)
		}
	}

	return result    
}

func main() {
	fmt.Println(transformArray([]int{4, 3, 2, 1}))
}
