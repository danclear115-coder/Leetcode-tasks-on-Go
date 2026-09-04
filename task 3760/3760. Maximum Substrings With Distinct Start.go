package main

import "fmt"

func maxDistinct(s string) int {

	seen := make(map[rune]struct{})
	
	for _, char := range s {
		seen[char] = struct{}{}
		fmt.Println(seen)
	}
	
	return len(seen)

}

func main() {
	fmt.Println(maxDistinct("abc"))
}
