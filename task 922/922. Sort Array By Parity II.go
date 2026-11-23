package main

import "fmt"

func initialCheck(nums []int) bool {

	for i := 0; i < len(nums); i++ {
		if (i % 2 == 0 && nums[i] % 2 != 0) || (i % 2 != 0 && nums[i] % 2 == 0) {
			return false
		}
	}

	return true

}

func getNums(nums []int) (even []int, notEven []int) {

	for _, num := range nums {
		if num % 2 == 0 {
			even = append(even, num)
		} else {
			notEven = append(notEven, num)
		}
	}

	return

}

func sortArrayByParityII(nums []int) []int {

	sortNums := []int{}
    
	if initialCheck(nums) {
		return nums
	}

	even, notEven := getNums(nums)

	indexEven, indexNotEven := 0, 0

	for i := 0; i < len(nums); i++ {
		if i % 2 == 0 {
			sortNums = append(sortNums, even[indexEven])
			indexEven++
		} else {
			sortNums = append(sortNums, notEven[indexNotEven])
			indexNotEven++
		}
		
	}

	return sortNums

}

func main() {
	fmt.Println(sortArrayByParityII([]int{4,2,5,7}))
}
