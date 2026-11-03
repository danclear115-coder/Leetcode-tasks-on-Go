package main

import "fmt"

func isInclude(word string, char byte) bool {

	for i := 0; i < len(word); i++ {
		if word[i] == char {
			return true
		}
	}

	return false

}

func findWordsContaining(words []string, x byte) []int {
    
	result := []int{}

	for i := 0; i < len(words); i++ {
		if isInclude(words[i], x) {
			result = append(result, i)
		}
	}

	return result

}

func main() {
	fmt.Println(findWordsContaining([]string{"abc","bcd","aaaa","cbc"}, 97))
}

