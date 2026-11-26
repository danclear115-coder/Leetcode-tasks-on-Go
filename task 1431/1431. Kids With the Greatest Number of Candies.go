package main

import "fmt"

func getMax(nums []int) int {

	max := nums[0]

	for _, number := range nums {
		if number > max {
			max = number
		}
	}

	return max

}

func kidsWithCandies(candies []int, extraCandies int) []bool {

	result := []bool{}

	maxCandies := getMax(candies)

	for _, numberOfCandies := range candies {
		if numberOfCandies + extraCandies < maxCandies {
			result = append(result, false)
		} else {
			result = append(result, true)
		}
	}

	return result

}

func main() {
	fmt.Println(kidsWithCandies([]int{2,3,5,1,3}, 3))
}
