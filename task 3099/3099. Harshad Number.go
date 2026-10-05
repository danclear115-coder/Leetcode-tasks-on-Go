package main

import "fmt"

func getSumOfNums(n int) (sum int) {

	for n > 0 {
		sum += n % 10
		n = n / 10
	}

	return

}

func sumOfTheDigitsOfHarshadNumber(x int) int {

	if x % getSumOfNums(x) == 0 {
		return getSumOfNums(x)
	}

	return -1
	
}

func main() {

	fmt.Println(sumOfTheDigitsOfHarshadNumber(18))

}
