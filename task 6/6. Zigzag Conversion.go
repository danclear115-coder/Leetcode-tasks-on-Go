package main

import "fmt"

func convert(s string, numRows int) string {
	if numRows == 1 || numRows >= len(s) {
		return s
	}

	rows := make([][]byte, numRows)

	row := 0
	step := 1

	for i := 0; i < len(s); i++ {
		rows[row] = append(rows[row], s[i])

		if row == 0 {
			step = 1
		} else if row == numRows-1 {
			step = -1
		}

		row += step
	}

	var result []byte

	for _, r := range rows {
		result = append(result, r...)
	}

	return string(result)
}


func main() {
	fmt.Println(convert("PAYPALISHIRING", 3))		
}
