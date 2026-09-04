package main

import "fmt"

func isInclude(str []string, element string) bool {

	for _, strElement := range str {
		if strElement == element {
			return true
		}
	}

	return false

}

func maxDistinct(s string) int {

	checkSlice := []string{}		

	for _, element := range s {
		if !isInclude(checkSlice, string(element)) {
			checkSlice = append(checkSlice, string(element))
		}
	}

	return len(checkSlice)

}

func main() {
	fmt.Println(maxDistinct("abab"))
}
