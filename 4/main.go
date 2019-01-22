package main

import (
	"fmt"
	"strconv"
)

func reverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

func isPalindromicNum(number int) bool {
	rnumber, _ := strconv.Atoi(reverse(strconv.Itoa(number)))
	return number == rnumber
}

func p4() int {
	var max = 0
	for i := 999; i > 99; i-- {
		for j := 999; j > 99; j-- {
			if isPalindromicNum(i * j) {
				// fmt.Println(i, j, i*j)
				if i*j > max {
					max = i * j
				}
			}
		}
	}
	return max
}

func main() {
	fmt.Println(p4())
}
