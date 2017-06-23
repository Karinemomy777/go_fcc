// Palindrome number algorithm
//
// 1. Get the number from user.
// 2. Reverse it.
// 3. Compare it with the number entered by the user.
// 4. If both are same then print palindrome number
// 5. Else print not a palindrome number.

package main

import (
	"fmt"
	"strings"
)

func main() {
	p1 := palindrome("eye")
	p2 := palindrome("Eye")
	p3 := palindrome("dog")
	fmt.Println(p1, p2, p3)
}

func palindrome(st string) string {
	s := strings.ToLower(st)
	r := reverseStrings(s)
	if s == r {
		return "This string is a palindrome! =) "
	}
	return "This string is not a palindrome! =( "

}
func reverseStrings(s string) string {
	splitting := []rune(s)
	for i, j := 0, len(splitting)-1; i < j; i, j = i+1, j-1 {
		splitting[i], splitting[j] = splitting[j], splitting[i]
	}
	return string(splitting)
}
