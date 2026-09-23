package main

import "fmt"

func containsDuplicate(nums []int) bool {
	visited := make(map[int]bool)

	for _, num := range nums {
		if visited[num] {
			return true
		}
		visited[num] = true
	}

	return false
}

func main() {
	fmt.Println(containsDuplicate([]int{1, 2, 3, 4, 1}))
}
