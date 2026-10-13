package main

import "fmt"

func getSumOfSquares(n int) int {

	sumOfSquares := 0
	
	for n > 0 {

		num := n % 10
		sumOfSquares += num * num
		n = n / 10

	}

	return sumOfSquares

}

func isHappy(n int) bool {

	for n != 1 && n != 4 {
		n = getSumOfSquares(n)
		if n == 4 {
			return false
		}
	}

	return true

}

func main() {
	fmt.Println(isHappy(2))
}
