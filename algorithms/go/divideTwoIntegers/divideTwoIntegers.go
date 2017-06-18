/*
 * Divide two integers without using multiplication, division and mod operator.
 * If it is overflow, return MAX_INT.
 */
package main

import (
	"fmt"
	"math"
)

func divide(dividend int, divisor int) int {
	result := 0
	flag := 0
	if divisor == 0 || dividend > math.MaxInt32 || dividend < math.MinInt32 || divisor > math.MaxInt32 || divisor < math.MinInt32 {
		return math.MaxInt32
	}

	if dividend == 0 {
		return 0
	}

	if (dividend > 0 && divisor < 0) || (dividend < 0 && divisor > 0) {
		flag = 1
	}

	if dividend < 0 {
		dividend = 0 - dividend
	}

	if divisor < 0 {
		divisor = 0 - divisor
	}

	for;dividend >= divisor;dividend = dividend - divisor {
		result = result + 1
	}

	if flag == 1 {
		result = 0 - result
	}

	if result > math.MaxInt32 {
		result = math.MaxInt32
	}

	return result
}

func main() {
	fmt.Println(divide(5,3))
	fmt.Println(divide(5,0))
	fmt.Println(divide(0,3))
	fmt.Println(divide(9,3))
	fmt.Println(divide(1,3))
	fmt.Println(divide(7,3))
	fmt.Println(divide(-7,3))
	fmt.Println(divide(9,-3))
	fmt.Println(divide(-2147483648, -1))
	fmt.Println(divide(-2147483648, 1))
}