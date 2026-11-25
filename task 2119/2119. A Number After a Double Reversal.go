package main

import (
	"strconv"
	"fmt"
)

func isSameAfterReversals(num int) bool {

    if num == 0 {
		return true
	}

	stringNum := strconv.Itoa(num)

    if stringNum[len(stringNum) - 1] == '0' {
        return false
    } else {
        return true
    }
    
}

func main() {
	fmt.Println(isSameAfterReversals(820047))
}
