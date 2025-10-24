package main

import (
	"fmt"
)

func nextBeautifulNumber(n int) int {
	for i := n + 1; i <= 1224444; i++ {
		if isBalance(i) {
			return i
		}
	}

	return -1
}

func isBalance(n int) bool {
	counter := make([]int, 10)

	for n > 0 {
		counter[n % 10]++
		n /= 10
	}

	for i := range 10 {
		if counter[i] > 0 && counter[i] != i {
			return false
		}
	}

	return true
}

func main() {
	// result: 22
	// n := int(1)

	// result: 1333
	// n := int(1000)

	// result: 3133
	n := int(3000)

	// result: 
	// n := int(0)

	result := nextBeautifulNumber(n)
	fmt.Printf("result = %v\n", result)
}

