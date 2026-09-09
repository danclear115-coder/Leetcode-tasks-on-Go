package main

import "fmt"

func getConcatenation(nums []int) []int {
    
	result := append(nums, nums...)
	return result

}

func main() {
	fmt.Println(getConcatenation([]int{1, 2, 3}))
}