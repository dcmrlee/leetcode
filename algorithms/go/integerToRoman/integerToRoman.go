/**********************************************************************************
*
* Given an integer, convert it to a roman numeral.
*
* Input is guaranteed to be within the range from 1 to 3999.
*
* Symbol:	 I	V	X	L	C	D	M
* Value :    1	5	10	50	100	500	1000
* Symbol:    IV	IX	XL	XC	CD	CM
* Value :    4	9	40	90	400	900
**********************************************************************************/

package main

import (
	"fmt"
)

func intToRoman(num int) string {
	symbols := []string{"M", "CM", "D", "CD", "C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"}
	values := []int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}
	var result string

	for i := 0; num != 0; i++ {
		for num >= values[i] {
			num -= values[i]
			result += symbols[i]
		}
	}

	return result
}

func main() {
	fmt.Printf("%s\n", intToRoman(10))
	fmt.Printf("%s\n", intToRoman(207))
	fmt.Printf("%s\n", intToRoman(1066))
}
