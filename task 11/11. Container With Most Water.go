package main

import "fmt"

func maxArea(height []int) int {
	left := 0
	right := len(height) - 1

	maxArea := 0

	for left < right {
		width := right - left

		h := height[left]
		if height[right] < h {
			h = height[right]
		}

		area := width * h

		if area > maxArea {
			maxArea = area
		}

		if height[left] < height[right] {
			left++
		} else {
			right--
		}
	}

	return maxArea
	
}

func main() {
	fmt.Println(maxArea([]int{1, 1}))
}
