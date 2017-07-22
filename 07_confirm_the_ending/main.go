// Check if a string (first argument, str) ends with the given
// target string (second argument, target).

package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	s := "Photograph"
	r := 'h'

	x := "Phonograph"
	y := 'p'

	confirmEnding(s, r)
	confirmEnding(x, y)
}

func confirmEnding(str string, target rune) {
	lenWord := len(str) - 1

	if strings.IndexRune(str, target) == lenWord {
		fmt.Printf("The string %s does contain the rune %s.\n", str, strconv.QuoteRune(target))
	} else {
		fmt.Printf("The string %s does NOT contain the rune %s.\n", str, strconv.QuoteRune(target))
	}
}
