package main

import "fmt"

func minimizedStringLength(s string) int {
	
	uniqueChars := make(map[byte]bool)

	for i := 0; i < len(s); i++ {
		uniqueChars[s[i]] = true
	}

	return len(uniqueChars)
}

func main() {
	fmt.Println(minimizedStringLength("aaabc"))
}
