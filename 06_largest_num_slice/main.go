// Return an array consisting of the largest number from each provided sub-array. For simplicity, the provided array will contain exactly 4 sub-arrays.
//
// Remember, you can iterate through an array with a simple for loop, and access each member with array syntax arr[i].

package main

import (
	"fmt"
)

var biggest, i int64

func main() {
	sl := []int64{24, 454, 1002, 12, 34, 764, 52}
	longestNum := iterateSlice(sl)
	fmt.Println("The longest number in the slice is:", longestNum)

}

func iterateSlice(s []int64) int64 {
	for _, v := range s {
		if v > i {
			i = v
			biggest = i
		}
	}
	return biggest
}
