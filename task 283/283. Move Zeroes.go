package main

func moveZeroes(nums []int) {

	writePointer := 0

	for a := 0; a < len(nums); a++ {
		if nums[a] != 0 {
			nums[writePointer] = nums[a]
			writePointer++
		}
	}

	for i := writePointer; i < len(nums); i++ {
		nums[i] = 0
	}

}

func main() {
	moveZeroes([]int{0, 1, 0})
}
