package main

import "fmt"

func main() {
	n := 100
	out := ((3 * n * n * n * n) + (2 * n * n * n) - (3 * n * n) - (2 * n)) / 12
	fmt.Println(out)
}
