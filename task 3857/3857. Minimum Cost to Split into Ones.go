package main

import "fmt"

func minCost(n int) int {

	return ( (n * (n - 1)) / 2)

}

func main() {
	fmt.Println(minCost(5))
}
