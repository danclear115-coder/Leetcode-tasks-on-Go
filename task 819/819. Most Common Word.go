package main

import (
	"fmt"
	"strings"
	"unicode"
)

type Word struct {
	word           string
	includeCounter int
}

func getUnbanned(words []Word, banned []string) (unbannedWords []string) {
	bannedMap := make(map[string]bool)
	for _, b := range banned {
		bannedMap[strings.ToLower(b)] = true
	}

	for _, w := range words {
		lowerWord := strings.ToLower(w.word)
		if !bannedMap[lowerWord] {
			unbannedWords = append(unbannedWords, lowerWord)
		}
	}

	return unbannedWords
}

func sortWords(words []Word) (sortedWords []Word) {
	tempWords := make([]Word, len(words))
	copy(tempWords, words)

	for len(tempWords) > 0 {
		maxWord, indexOfMax := tempWords[0], 0

		for i := 0; i < len(tempWords); i++ {
			if tempWords[i].includeCounter > maxWord.includeCounter {
				maxWord, indexOfMax = tempWords[i], i
			}
		}

		sortedWords = append(sortedWords, maxWord)
		tempWords = append(tempWords[:indexOfMax], tempWords[indexOfMax+1:]...)
	}

	return sortedWords
}

func mostCommonWord(paragraph string, banned []string) string {
	paragraph = strings.Map(func(r rune) rune {
		if unicode.IsLetter(r) {
			return r
		}
		return ' '
	}, paragraph)

	mostRepeatedWords, words := []Word{}, strings.Fields(paragraph)

	for a := 0; a < len(words); a++ {
		counter := 1
		for b := 0; b < len(words); b++ {
			if a != b && strings.EqualFold(words[a], words[b]) {
				counter += 1
			}
		}
		mostRepeatedWords = append(mostRepeatedWords, Word{word: words[a], includeCounter: counter})
	}

	unbannedWords := getUnbanned(sortWords(mostRepeatedWords), banned)

	if len(unbannedWords) > 0 {
		return unbannedWords[0]
	} else {
		return ""
	}
}

func main() {
	fmt.Println(mostCommonWord("a.", []string{}))
}
