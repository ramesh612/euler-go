package main

import (
	"fmt"
	"math"
)

func triangleNumber(n int) int {
	return n * (n + 1) / 2
}

func divisors(n int) []int {
	var out []int
	for i := 1; i <= int(math.Sqrt(float64(n))); i++ {
		if n%i == 0 {
			if n/i == i {
				out = append(out, i)
			} else {
				out = append(out, i)
				out = append(out, n/i)
			}
		}
	}
	return out
}

func main() {
	for i := 1; ; i++ {
		N := triangleNumber(i)
		divisorCount := len(divisors(N))
		if divisorCount > 500 {
			fmt.Println(N, divisorCount)
			return
		}
	}
}
