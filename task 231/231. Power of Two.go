package main

import "fmt"

func isPowerOfTwo(n int) bool {

    if n == 1 {
		return true
	} else if n == 0 {
        return false
    }

	for i := 0; i < 62; i++ {
		if 2 << i == n {
			return true
		}
	}

	return false
}

func main() {
	fmt.Println(isPowerOfTwo(1))
}