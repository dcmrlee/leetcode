/**********************************************************************************
*
* The string "PAYPALISHIRING" is written in a zigzag pattern on a given number of rows like this:
* (you may want to display this pattern in a fixed font for better legibility)
*
* P   A   H   N
* A P L S I I G
* Y   I   R
*
* And then read line by line: "PAHNAPLSIIGYIR"
*
* Write the code that will take a string and make this conversion given a number of rows:
*
* string convert(string text, int nRows);
*
* convert("PAYPALISHIRING", 3) should return "PAHNAPLSIIGYIR".
*
*
*
* 1     7      13
* 2   6 8    12
* 3 5   9  11
* 4     10
*
**********************************************************************************/
package main

import "fmt"

func convert(s string, numRows int) string {
	str_len := len(s)

	if str_len == 0 || numRows <= 1 || numRows >= str_len {
		return s
	}

	r := make(map[int]string, str_len)
	row := 0
	step := 1
	var result string

	for i := 0; i < str_len; i++ {
		if row == numRows-1 {
			step = -1
		}
		if row == 0 {
			step = 1
		}
		r[row] += string(s[i])
		row += step
	}

	for i := 0; i < numRows; i++ {
		result += r[i]
	}
	return string(result)
}

func main() {
	str1 := "PAYPALISHIRING"
	fmt.Println(convert(str1, 3))
}
