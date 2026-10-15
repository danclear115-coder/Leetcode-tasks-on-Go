package main

import "fmt"

func smallestEvenMultiple(n int) int {

	for i := n; i <= n*2; i++ {
		if i%2 == 0 && i%n == 0 {
			return i
		}
	}

	return 0

}

func main() {
	fmt.Println(smallestEvenMultiple(6))
}
