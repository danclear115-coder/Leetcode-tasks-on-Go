package main

import "fmt"

func reverse(x int) int {

	result, isMin := 0, false

	if x < 0 {
		x = x * -1
		isMin = true
	}

	for x > 0 {
		num := x % 10
		result = result*10 + num
		x = x / 10
	}

	if result > 2147483647 || result < -2147483648 {
		return 0
	}

	if isMin {
		return result * -1
	} else {
		return result
	}

}

func main() {
	fmt.Println(reverse(-123))
}
