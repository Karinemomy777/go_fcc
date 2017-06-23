package main

import (
	"fmt"
)

func main() {
	//reverseStrings("Golang")
	s := reverseStrings("Golang")
	fmt.Println(s)
}

func reverseStrings(s string) string {
	splitting := []rune(s)
	for i, j := 0, len(splitting)-1; i < j; i, j = i+1, j-1 {
		splitting[i], splitting[j] = splitting[j], splitting[i]
	}
	return string(splitting)
}
