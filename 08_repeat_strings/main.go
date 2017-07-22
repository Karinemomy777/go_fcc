// Repeat a given string (first argument) num times (second argument).
// Return an empty string if num is not a positive number.

package main

import (
	"fmt"
	"strings"
)

func main() {
	s := "This is an amazing string!"
	fmt.Println(strings.Repeat(s, 3))
}

func repeatString(s string, n int) string {
	sl := make([]byte, len(s)*n)
	c := copy(sl, s)
	for c < len(sl) {
		copy(sl[c:], sl[:c])
		c *= 2
	}
	return string(sl)
}
