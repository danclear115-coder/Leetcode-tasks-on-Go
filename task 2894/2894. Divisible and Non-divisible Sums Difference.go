package main

import "fmt"

func differenceOfSums(n int, m int) int {

	sumOfDiv, sumOfNonDiv := 0, 0

	for i := 1; i <= n; i++ {
		if i % m == 0 {
			sumOfDiv += i
		} else {
			sumOfNonDiv += i
		}
	}

	return sumOfNonDiv - sumOfDiv

}

func main() {
	fmt.Println(differenceOfSums(10, 3))
}
