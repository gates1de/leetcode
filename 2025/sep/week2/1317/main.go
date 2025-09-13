package main

import (
	"fmt"
	"strconv"
	"strings"
)

func getNoZeroIntegers(n int) []int {
	for a := 1; a < n; a++ {
		b := n - a
		if !strings.Contains(strconv.Itoa(a), "0") && !strings.Contains(strconv.Itoa(b), "0") {
				return []int{a, b}
		}
	}

	return []int{}
}

func main() {
	// result: [1,1]
	// n := int(2)

	// result: [2,9]
	n := int(11)

	// result: []
	// n := int(0)

	result := getNoZeroIntegers(n)
	fmt.Printf("result = %v\n", result)
}

