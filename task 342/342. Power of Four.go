package main

import "fmt"

func isPowerOfFour(n int) bool {

	for n%4 == 0 {
		n /= 4
	}

	return n == 1

}

func main() {
	fmt.Println(isPowerOfFour(1))
}
