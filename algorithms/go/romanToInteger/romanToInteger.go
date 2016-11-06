/**********************************************************************************
*
* Given a roman numeral, convert it to an integer.
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

func romanToInt(s string) int {
	symbolToValue := map[byte]int{
		'M': 1000,
		'D': 500,
		'C': 100,
		'L': 50,
		'X': 10,
		'V': 5,
		'I': 1,
	}
	result := 0
	preValue := 0
	curValue := 0
	for i := 0; i < len(s); i++ {
		curValue = symbolToValue[s[i]]
		if curValue <= preValue {
			result += preValue
		} else {
			result -= preValue
		}
		preValue = curValue
	}

	result += curValue

	return result
}

func main() {
	fmt.Printf("%d\n", romanToInt("DCXXI"))
	fmt.Printf("%d\n", romanToInt("IX"))
	fmt.Printf("%d\n", romanToInt("CD"))
}
