/**********************************************************************************
*
* Implement atoi to convert a string to an integer.
*
* Hint: Carefully consider all possible input cases. If you want a challenge,
*       please do not see below and ask yourself what are the possible input cases.
*
* Notes:
*   It is intended for this problem to be specified vaguely (ie, no given input specs).
*   You are responsible to gather all the input requirements up front.
*
*
* Requirements for atoi:
*
* The function first discards as many whitespace characters as necessary until the first
* non-whitespace character is found. Then, starting from this character, takes an optional
* initial plus or minus sign followed by as many numerical digits as possible, and interprets
* them as a numerical value.
*
* The string can contain additional characters after those that form the integral number,
* which are ignored and have no effect on the behavior of this function.
*
* If the first sequence of non-whitespace characters in str is not a valid integral number,
* or if no such sequence exists because either str is empty or it contains only whitespace
* characters, no conversion is performed.
*
* If no valid conversion could be performed, a zero value is returned. If the correct value
* is out of the range of representable values, INT_MAX (2147483647) or INT_MIN (-2147483648)
* is returned.
*
**********************************************************************************/
package main

import "fmt"
import "math"

func myAtoi(str string) int {
	bytes := []byte(str)
	ret := 0
	flag := 1
	i := 0

	if len(bytes) == 0 {
		return 0
	}

	// filter space
	for ; i < len(bytes); i++ {
		if bytes[i] == ' ' {
			continue
		} else {
			break
		}
	}

	// negative or positive
	if bytes[i] == '-' || bytes[i] == '+' {
		if bytes[i] == '-' {
			flag = -1
		}
		i++
	}

	for ; i < len(bytes); i++ {
		if bytes[i] > '9' || bytes[i] < '0' {
			break
		}
		digit := int(bytes[i] - '0')
		if flag == -1 {
			if -ret < (math.MinInt32+digit)/10 {
				return math.MinInt32
			}
		} else {
			if ret > (math.MaxInt32-digit)/10 {
				return math.MaxInt32
			}
		}
		ret = ret*10 + digit
	}
	return ret * flag
}

func main() {
	fmt.Println(myAtoi("     +004500"))
	fmt.Println(myAtoi("+-2"))
	fmt.Println(myAtoi(""))
	fmt.Println(myAtoi("  -0012a42"))
}
