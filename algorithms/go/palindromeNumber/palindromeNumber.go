package main

import (
	"fmt"
)

func isPalindrome(x int) bool {
	tmp := x
	newx := 0
	for {
		if tmp <= 0 {
			break
		}
		last_digit := tmp % 10
		tmp /= 10
		newx = newx*10 + last_digit
	}
	if newx == x {
		return true
	}
	return false
}

func main() {
	fmt.Println(isPalindrome(0))
	fmt.Println(isPalindrome(-101))
	fmt.Println(isPalindrome(1001))
	fmt.Println(isPalindrome(1234321))
	fmt.Println(isPalindrome(2147447412))
	fmt.Println(isPalindrome(2142))
}
