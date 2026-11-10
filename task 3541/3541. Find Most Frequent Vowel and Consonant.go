package main

import "fmt"

func getSymbolIncludeCounter(s string, symbol byte) (counter int) {

	for i := 0; i < len(s); i++ {
		if s[i] == symbol {
			counter++
		}
	}

	return

}

func isInclde(s []string, symbol string) bool {

	for i := 0; i < len(s); i++ {
		if s[i] == symbol {
			return true
		}
	}

	return false

}

func maxFreqSum(s string) int {

	vowelsCounter, consonantsCounter := 0, 0

	vowels := []string{"a", "e", "i", "o", "u"}

	consonants := []string{
		"b", "y", "c", "d", "f", "g", "h", "j", "k", "l", "m",
		"n", "p", "q", "r", "s", "t", "v", "w", "x", "z",
	}

	for i := 0; i < len(s); i++ {
		if isInclde(vowels, string(s[i])) && vowelsCounter < getSymbolIncludeCounter(s, s[i]) {
			vowelsCounter = getSymbolIncludeCounter(s, s[i])
		} else if isInclde(consonants, string(s[i])) && consonantsCounter < getSymbolIncludeCounter(s, s[i]) {
			consonantsCounter = getSymbolIncludeCounter(s, s[i])
		}
		fmt.Println(s[i], vowelsCounter, consonantsCounter)
	}

	fmt.Println(vowelsCounter, consonantsCounter)
	return vowelsCounter + consonantsCounter
}

func main() {
	fmt.Println(maxFreqSum("yj"))
}
