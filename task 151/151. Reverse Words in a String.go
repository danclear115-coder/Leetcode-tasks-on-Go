package main

import (
	"fmt"
	"strings"
)

func reverseWords(s string) string {

	reversedString := ""

	wordsSlice := strings.Fields(s)

	for i := len(wordsSlice) - 1; i >= 0; i-- {
		reversedString += wordsSlice[i]
		if i != 0 {
			reversedString += " "
		}
	}

	return reversedString
    
}

func main() {
	fmt.Println(reverseWords("a good   example"))
}
