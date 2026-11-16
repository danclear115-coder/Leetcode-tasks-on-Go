package main

import (
	"fmt"
	"strconv"
)

func GetMaxAndUpdate(num int) (int, int) {

	numSlice := []int{}

	for num > 0 {
		numSlice = append(numSlice, num%10)
		num = num / 10
	}

	max, index := 0, 0 // a number less than zero is not possible in this example.

	for i := 0; i < len(numSlice); i++ {
		if numSlice[i] > max {
			max, index = numSlice[i], i
		}
	}

	strNewNum := ""

	for i := 0; i < len(numSlice); i++ {
		if i != index {
			strNum := strconv.Itoa(numSlice[i])
			strNewNum += strNum
		}
	}

	newNum, err := strconv.Atoi(strNewNum)

	if err != nil {

	}

	return max, newNum

}

func maxProduct(n int) int {

	max1, newNum1 := GetMaxAndUpdate(n)
	max2, _ := GetMaxAndUpdate(newNum1)

	return max1 * max2

}

func main() {
	fmt.Println(maxProduct(12435346971293))
	fmt.Println(GetMaxAndUpdate(12435346971293))
}
