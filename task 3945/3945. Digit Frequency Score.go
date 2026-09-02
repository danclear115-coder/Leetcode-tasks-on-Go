package main

import (
	"fmt"
	"strconv"
)

func isInclude(nums []int, num int) bool {
	
	for _, number := range nums {
		if number == num {
			return true
		}
	}

	return false

}

func digitFrequencyScore(n int) int {
	
	sum, stringNum, includeSlice := 0, strconv.Itoa(n), []int{}

	for a := 0; a < len(stringNum); a++ {
		
		counter := 1;

		for b := 0; b < len(stringNum); b++ {
			if stringNum[a] == stringNum[b] && a != b {
				counter++
			}
		}

		number, err := strconv.Atoi(string(stringNum[a]))

		if err == nil && !isInclude(includeSlice,  number) {
			includeSlice = append(includeSlice, number)
			sum += number * counter
		}
				
	}

	return sum

}

func main() {
	fmt.Println(digitFrequencyScore(101))
}
