package main

import (
	"fmt"
	"strconv"
)

func reverseStrNum(strNum string) (reverseStrNum string) {

	for i := len(strNum) - 1; i >= 0; i-- {
		reverseStrNum += string(strNum[i])
	}

	return

}

func mirrorDistance(n int) int {
    
	strNum := strconv.Itoa(n)
	reverseStrNum := reverseStrNum(strNum)
	reverseNum, err := strconv.Atoi(reverseStrNum)

	fmt.Println(reverseNum, n)
	
	if err == nil {
        if n - reverseNum < 1 {
            return (n - reverseNum) * -1
        } else {
            return n - reverseNum
        }
	}

	result := 0

	return result

}

func main() {
	fmt.Println(mirrorDistance(10))
}
