package main

import "fmt"

func longestPalindrome(s string) string {
	
	if len(s) < 2 {
		return s
	}

	start, maxLen := 0, 0

	for i := 0; i < len(s); i++ {

		len1 := expandAroundCenter(s, i, i)
		len2 := expandAroundCenter(s, i, i+1)

		currentLen := len1
		if len2 > len1 {
			currentLen = len2
		}

		if currentLen > maxLen {
			maxLen = currentLen
			start = i - (currentLen-1)/2
		}
	}

	return s[start : start+maxLen]
}

func expandAroundCenter(s string, left, right int) int {
	for left >= 0 && right < len(s) && s[left] == s[right] {
		left--
		right++
	}
	return right - left - 1
}


func main() {
	fmt.Println(longestPalindrome("babad"))
}
