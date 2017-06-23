package main

import (
	"fmt"
	"strings"
)

func main() {
	word := longestWord("The Go Programming Language")
	fmt.Println(word)
}

func longestWord(s string) int {
	st := strings.Fields(s)
	longest := 1

	for i := 0; i < len(st); i++ {
		if longest < len(st[i]) {
			longest = len(st[i])
		}
	}
	return longest
}

//Return longest word in a string.
