package main

import (
	"fmt"
)

func main() {
	num := 8

	f := factorialize(num)

	//fmt.Println("The factorial of " + string(num) + " is " + string(f) + ".")
	fmt.Println(<-f)
}

func factorialize(n int) <-chan int {
	start := make(chan int)
	go func() {
		total := 1
		for i := n; i > 0; i-- {
			total *= i
		}
		start <- total
		close(start)
	}()

	return start
}
