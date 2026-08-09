package main

import "fmt"

func lengthOfLongestSubstring(s string) int {

	if len(s) == 0 {
		return 0
	}

	maxLen := 0
	left := 0
	charIndexMap := make(map[byte]int)

	for right := 0; right < len(s); right++ {

		if index, found := charIndexMap[s[right]]; found && index >= left {
			left = index + 1
		}
		
		charIndexMap[s[right]] = right
		
		currentLen := right - left + 1
		
		if currentLen > maxLen {
			maxLen = currentLen
		}
	}

	return maxLen
}

func main() {
	fmt.Println(lengthOfLongestSubstring("abcabcbb"))
}
