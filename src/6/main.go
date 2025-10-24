package main

import (
	"fmt"
)

func main() {
	res := convert("PAYPALISHIRING", 3)
	fmt.Println(res)
}

func convert(s string, numRows int) string {
	if numRows == 1 {
		return s
	}

	res := ""
	for row := 0; row < numRows; row++ {
		flag := true
		for i := row; i < len(s); {
			res += s[i : i+1]
			if row == 0 || row == numRows-1 {
				i += (numRows - 1) * 2
				continue
			}

			if flag {
				i += (numRows - row - 1) * 2
				flag = false
			} else {
				i += row * 2
				flag = true
			}
		}
	}
	return res
}
