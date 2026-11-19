package main

import "fmt"

func removeDuplicateLetters(s string) string {
	lastPos := make(map[byte]int)
	for i := 0; i < len(s); i++ {
		lastPos[s[i]] = i
	}

	stack := []byte{}
	inStack := make(map[byte]bool)

	for i := 0; i < len(s); i++ {
		
		ch := s[i]

		if inStack[ch] {
			continue
		}

		for len(stack) > 0 && stack[len(stack)-1] > ch && lastPos[stack[len(stack)-1]] > i {
			top := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			inStack[top] = false
		}

		stack = append(stack, ch)
		inStack[ch] = true
	}

	return string(stack)
}

func main() {
	fmt.Println(removeDuplicateLetters("cbacdcbc"))
}
