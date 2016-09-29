/**********************************************************************************
*
* Implement regular expression matching with support for '.' and '*'.\
*
* '.' Matches any single character.
* '*' Matches zero or more of the preceding element.

* The matching should cover the entire input string (not partial).

* The function prototype should be:
* bool isMatch(const char *s, const char *p)

* Some examples:
* isMatch("aa","a") → false
* isMatch("aa","aa") → true
* isMatch("aaa","aa") → false
* isMatch("aa", "a*") → true
* isMatch("aa", ".*") → true
* isMatch("ab", ".*") → true
* isMatch("aab", "c*a*b") → true
* isMatch("ab", ".*c") → false
**********************************************************************************/
package main

import "fmt"

func isMatch(s string, p string) bool {
	/*

		 This is a recursive method.
		 You start to “consume” p and s until p can be consumed to empty.
		 If at this moment s happens to be consumed to empty also,
		 your match is successful.

		 Everytime you look at the first and the second character in p.
		 There are two situations:
		   - If you find the second character in p is “*”.
		     This means, on the left of “*” in p, there is another character, x.
				a) x is not . and the first position of s is not the same as x.
				   Then you must skip x* in p to  continue to match s.
				   So you rely on isMatch(s, p[2:])
				b) x is not . and the first position of s is the same as x.
				   Then whether p matches s will depend on ‘isMatch(s[1:], p)’ or isMatch(s, p[2:])
				c) x is .. Then whether p matches s will also depend on isMatch(s[1:], p)
		   - If you find the second character in p is not “*”,
		     then a successful match will only happen if the first character of p and s
			 are the same and the rest of p and s match

	*/

	if len(p) == 0 {
		return len(s) == 0
	}

	if len(p) >= 2 && p[1] == '*' {
		return len(s) != 0 && (s[0] == p[0] || p[0] == '.') && isMatch(s[1:], p) || isMatch(s, p[2:])
	} else {
		return len(s) != 0 && (s[0] == p[0] || p[0] == '.') && isMatch(s[1:], p[1:])
	}
	return false
}

func main() {
	fmt.Println(isMatch("aa", "a"))
	fmt.Println(isMatch("aa", "aa"))
	fmt.Println(isMatch("aaa", "aa"))
	fmt.Println(isMatch("aa", "a*"))
	fmt.Println(isMatch("aa", ".*"))
	fmt.Println(isMatch("ab", ".*"))
	fmt.Println(isMatch("aab", "c*a*b"))
	fmt.Println(isMatch("ab", ".*c"))
}
