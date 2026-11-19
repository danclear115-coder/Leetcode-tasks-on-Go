package main

import "fmt"

func isUgly(n int) bool {

	if n <= 0 {
		return false
	}

	divs := []int{2, 3, 5}

	for _, div := range divs {
		for n % div == 0 {
			n /= div 
		}
	}

	return n == 1

}

func main() {
	fmt.Println(isUgly(14))
}
