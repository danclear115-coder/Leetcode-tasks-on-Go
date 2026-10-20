package main

import "fmt"

func getNewNum(num int) (newNum int) {

	for num > 0 {
		newNum += num % 10
		num = num / 10
	}

	return newNum

}

func addDigits(num int) int {
    
	for num > 9 {
		num = getNewNum(num)
	}

	return num

}

func main() {
	fmt.Println(addDigits(38))
	
}
