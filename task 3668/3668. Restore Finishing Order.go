package main

import "fmt"

func recoverOrder(order []int, friends []int) []int {

	result := []int{}

	for a := 0; a < len(order); a++ {
		for b := 0; b < len(friends); b++ {
			if order[a] == friends[b] {
				result = append(result, order[a])
				break
			}
		}
	}

	return result

}

func main() {
	fmt.Println(recoverOrder([]int{3, 1, 2, 5, 4}, []int{1, 3, 4}))
}
