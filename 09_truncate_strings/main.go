package main

import (
	"fmt"
)

func main() {
	v := truncateString("A-tisket a-tasket A green and yellow basket", 11)
	fmt.Println(v)
}

// Truncate the string if the num of char in the string is larger than the num argument

func truncateString(str string, num int) string {
	trunc := str

	if len(str) > num {
		if num > 3 {
			num -= 3
		}
		trunc = str[0:num] + "..."
	}
	return trunc

}
