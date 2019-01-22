package main

import "fmt"

func makeRange(min, max int) []int {
	a := make([]int, max-min+1)
	for i := range a {
		a[i] = min + i
	}
	return a
}

// All returns true if all of the integers in the slice satisfy the predicate f.
func All(vs []int, f func(int, int) bool, value int) bool {
	for _, v := range vs {
		if !f(v, value) {
			return false
		}
	}
	return true
}

// Any returns true if one of the integers in the slice satisfies the predicate f.
func Any(vs []int, f func(int, int) bool, value int) bool {
	for _, v := range vs {
		if f(v, value) {
			return true
		}
	}
	return false
}

func allEqualTo(v int, value int) bool {
	return v == value
}

func allDivides(v int, value int) bool {
	return v%value == 0
}

func divide(a []int, number int) []int {
	for i, v := range a {
		if v%number == 0 {
			a[i] = v / number
		}
	}
	return a
}

func main() {
	max := 20
	a := makeRange(2, max)

	var output = 1
	for i := 2; i <= max; i++ {
		for Any(a, allDivides, i) {
			fmt.Println(i, divide(a, i))
			output *= i
		}
	}

	fmt.Println(output, All(a, allEqualTo, 1))
}
