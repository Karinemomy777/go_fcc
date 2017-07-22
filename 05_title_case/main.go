package main

import (
	"fmt"
	"strings"
)

func main() {
	str := titleString("i really love learning golang")

	fmt.Println(str)
}

func titleString(s string) string {
	words := strings.Fields(s)

	for index, word := range words {
		if words[index] == 1 {
			words[index] = strings.Title(word)
		} else {
			words[index] = word
		}

	}
	return strings.Join(words, " ")
}
