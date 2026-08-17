package main

import (
	"fmt"
	"math"
)

func myAtoi(s string) int {
	n := len(s)
	if n == 0 {
		return 0
	}

	i := 0

	for i < n && s[i] == ' ' {
		i++
	}

	if i == n {
		return 0
	}

	sign := 1
	if s[i] == '+' {
		i++
	} else if s[i] == '-' {
		sign = -1
		i++
	}

	result := 0

	maxInt32 := math.MaxInt32
	minInt32 := math.MinInt32

	for i < n && s[i] >= '0' && s[i] <= '9' {
		digit := int(s[i] - '0')

		if result > maxInt32/10 || (result == maxInt32/10 && digit > maxInt32%10) {
			if sign == 1 {
				return maxInt32
			}
			return minInt32
		}

		result = result*10 + digit
		i++
	}

	return result * sign
}

func main() {
	fmt.Println(myAtoi("1337c0d3"))
}
