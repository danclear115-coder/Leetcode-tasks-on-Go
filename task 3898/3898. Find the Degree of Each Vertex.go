package main

import "fmt"

func findDegrees(matrix [][]int) []int {
    
	result := []int{}

	for _, slice := range matrix {

		num := 0

		for i := 0; i < len(slice); i++ {
			if slice[i] != 0 {
				num += slice[i]
			}
		}

		result, num = append(result, num), 0

	}

	return result

}

func main() {
	fmt.Println(findDegrees([][]int{{0}}))
}
