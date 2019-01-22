package main

import "fmt"

func divides(number int, divisor int) bool {
	if number%divisor == 0 {
		return true
	}
	return false
}

func main() {
	var sum int
	for i := 1; i < 1000; i++ {
		if divides(i, 3) || divides(i, 5) {
			sum += i
		}
	}
	fmt.Println(sum)
}
