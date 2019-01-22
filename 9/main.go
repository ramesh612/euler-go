package main

import (
	"fmt"
	"math"
)

func main() {
	for n := 1; ; n++ {
		for m := 1; m < n; m++ {
			a := int(math.Abs(float64((m * m) - (n * n))))
			b := 2 * m * n
			c := (m * m) + (n * n)
			if a+b+c == 1000 {
				fmt.Printf("a == %d, b == %d, c == %d, product == %d\n", a, b, c, a*b*c)
				return
			} else if a+b+c > 1000 {
				break
			}
		}
	}
}
