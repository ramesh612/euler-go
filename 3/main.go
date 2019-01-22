package main

import (
	"fmt"
	"math"
)

func primeFactors(n int) {
	// Print the number of 2s that divide n
	for n%2 == 0 {
		fmt.Printf("%d ", 2)
		n = n / 2
	}

	// n must be odd at this point.  So we can skip
	// one element (Note i = i +2)
	for i := 3; i <= int(math.Sqrt(float64(n))); i = i + 2 {
		// While i divides n, print i and divide n
		for n%i == 0 {
			fmt.Printf("%d ", i)
			n = n / i
		}
	}

	// This condition is to handle the case when n
	// is a prime number greater than 2
	if n > 2 {
		fmt.Printf("%d ", n)
	}
}

func main() {
	primeFactors(600851475143)
}
