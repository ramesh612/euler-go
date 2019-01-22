package main

import "fmt"

func evenFibSum(limit int) int {
	var total int
	var previous = 1
	var current = 2
	for current < limit {
		var temp = current
		current += previous
		previous = temp
		fmt.Printf("next fib is: %d\n", current)
		if current%2 == 0 {
			total += current
		}
	}
	return total + 2 // because we are starting from 3, we have to add the first even term 2.
}

func main() {
	fmt.Println(evenFibSum(4000000))
}
