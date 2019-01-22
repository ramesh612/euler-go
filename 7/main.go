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

func printNPrimes(N int) {
	primeCount := 0
	for number := 2; primeCount < N; number++ {
		if isPrime(number) {
			fmt.Printf("%d ", number)
			primeCount++
		}
	}
	fmt.Println()
}

func main() {
	n, _ := strconv.Atoi(os.Args[1])
	printNPrimes(n)
}
