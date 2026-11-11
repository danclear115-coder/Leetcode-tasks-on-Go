package main

import (
	"fmt"
	"strconv"
)

func getPicksAndValley(num int) (counter int) {

	stringNum := strconv.Itoa(num)

	for i := 1; i < len(stringNum) - 1; i++ {

		prevNum, err1 := strconv.Atoi(string(stringNum[i - 1]))
		actualNum, err2 := strconv.Atoi(string(stringNum[i]))
		nextNum, err3 := strconv.Atoi(string(stringNum[i + 1]))

		if err1 == nil && err2 == nil && err3 == nil {
			//fmt.Println("Error")
		}

		if (actualNum < prevNum && actualNum < nextNum) || (actualNum > prevNum && actualNum > nextNum) {
			counter++
		}

	}

	return

}

func totalWaviness(num1 int, num2 int) int {
    
	counter := 0

	for i := num1; i <= num2; i++ {
		counter += getPicksAndValley(i)
	}

	return counter

}

func main() {
	fmt.Println(totalWaviness(120, 130))
}
