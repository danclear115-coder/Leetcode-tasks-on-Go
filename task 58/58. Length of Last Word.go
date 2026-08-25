package main

import "fmt"

func lengthOfLastWord(s string) int {

	counter := 0

	if len(s) == 1 {
        return 1
    }

	for i := len(s) - 1; i >= 0; i-- {
		fmt.Println(string(s[i]))
		if counter == 0 && string(s[i]) != " " {
			counter++
		} else if counter > 0 && string(s[i]) != " " {
			counter++
		}else if counter > 0 && string(s[i]) == " " {
			break
		}
	}

	return counter

}

func main() {
	fmt.Println(lengthOfLastWord("a "))
}
