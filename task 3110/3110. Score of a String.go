package main

import "fmt"

func inModule(n int) int {

	if n < 0 {
		return n * -1
	}

	return n

} 

func scoreOfString(s string) int {

	sum := 0

	for i := 0; i < len(s); i++ {
		if i + 1 < len(s) {
			num1, num2 := int(s[i]), int(s[i + 1])
			sum += inModule((num1 - num2))
		}
	}

	return sum

}

func main() {
	fmt.Println(scoreOfString("hello"))
}
