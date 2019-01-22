package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
)

func isPrime(n int) bool {
	for i := 2; i <= int(math.Sqrt(float64(n))); i++ {
		if n%i == 0 {
			return false
		}
	}

	return true
}

func sumOfPrimesLessThan(N int) int {
	sum := 0
	for number := 2; number < N; number++ {
		if isPrime(number) {
			sum += number
		}
	}
	return sum
}

func main() {
	n, _ := strconv.Atoi(os.Args[1])
	fmt.Println(sumOfPrimesLessThan(n))
}
