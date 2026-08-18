package main

import (
	"fmt"
	"strconv"
)

func reverseString(str string) (revStr string) {

	for i := len(str) - 1; i >= 0; i-- {
		revStr += string(str[i])
	}

	return

}

func isPalindrome(x int) bool {

	str := strconv.Itoa(x)
	reverseStr := reverseString(str)

	return str == reverseStr

}

func main() {
	
	fmt.Println(isPalindrome(10))
}
