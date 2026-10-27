package main

import (
	"fmt"
)

type LetterWeight struct {
	Char   string
	Weight int
}

func getReversedAlphabet() []LetterWeight {

	reversedAlphabet := make([]LetterWeight, 26)

	for i := 0; i < 26; i++ {
		char := rune('a' + i)
		reversedAlphabet[i] = LetterWeight{
			Char:   string(char),
			Weight: 26 - i,
		}
	}

	return reversedAlphabet

}

func reverseDegree(s string) int {

	sum, reversedAlphabet := 0, getReversedAlphabet()

	for a := 0; a < len(s); a++ {
		for b := 0; b < len(reversedAlphabet); b++ {
			if string(s[a]) == reversedAlphabet[b].Char {
				sum += reversedAlphabet[b].Weight * (a + 1)
			}
		}
	}

	return sum

}

func main() {
	fmt.Println(reverseDegree("abc"))
}
