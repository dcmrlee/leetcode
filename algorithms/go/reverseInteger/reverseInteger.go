/**********************************************************************************
*
* Reverse digits of an integer.
*
* Example1: x =  123, return  321
* Example2: x = -123, return -321
*
*
* Have you thought about this?
*
* Here are some good questions to ask before coding. Bonus points for you if you have already thought through this!
*
* > If the integer's last digit is 0, what should the output be? ie, cases such as 10, 100.
*
* > Did you notice that the reversed integer might overflow? Assume the input is a 32-bit integer,
*   then the reverse of 1000000003 overflows. How should you handle such cases?
*
* > Throw an exception? Good, but what if throwing an exception is not an option?
*   You would then have to re-design the function (ie, add an extra parameter).
*
*
**********************************************************************************/
package main

import (
	"fmt"
	"math"
)

func reverse(x int) int {
	// math.MaxInt32 == 2147483647 == 1<<31 -1
	// math.MinInt32 == -2147483648 == -1<<31
	var y int
	var n int

	for {
		if x == 0 {
			break
		}
		n = x % 10
		if y > math.MaxInt32/10 || y < math.MinInt32/10 {
			return 0
		}
		y = y*10 + n
		x /= 10
	}
	return y
}

func main() {
	fmt.Println(math.MaxInt32)
	fmt.Println(math.MinInt32)
	fmt.Println(reverse(123))
	fmt.Println(reverse(1463847412))
	fmt.Println(reverse(1000000003))
	fmt.Println(reverse(-2147483648))
}
