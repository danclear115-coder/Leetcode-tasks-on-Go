package main

import "fmt"

func isMatch(s string, p string) bool {
	
	m := len(s)
	n := len(p)

	dp := make([][]bool, m+1)

	for i := 0; i <= m; i++ {
		dp[i] = make([]bool, n+1)
	}

	dp[0][0] = true

	for j := 2; j <= n; j++ {
		if p[j-1] == '*' {
			dp[0][j] = dp[0][j-2]
		}
	}

	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if p[j-1] == s[i-1] || p[j-1] == '.' {
				dp[i][j] = dp[i-1][j-1]
			}
			if p[j-1] == '*' {
				dp[i][j] = dp[i][j-2]
				if p[j-2] == s[i-1] || p[j-2] == '.' {
					dp[i][j] = dp[i][j] || dp[i-1][j]
				}
			}
		}
	}

	return dp[m][n]
}


func main() {
	fmt.Println(isMatch("ab", "*"))
}
