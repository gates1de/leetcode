package main
import (
	"fmt"
)

func equalSubstring(s string, t string, maxCost int) int {
	n := len(s)
	result := int(0)
	start := int(0)
	currentCost := int(0)

	for i := 0; i < n; i++ {
		currentCost += abs(int(s[i]) - int(t[i]))

		for currentCost > maxCost {
			currentCost -= abs(int(s[start]) - int(t[start]))
			start++
		}

		result = max(result, i - start + 1)
	}

	return result
}

func abs(num int) int {
	if num < 0 {
		return -num
	}
	return num
}

func max(a, b int) int {
	if b > a {
		return b
	}
	return a
}

func main() {
	// result: 3
	// s := "abcd"
	// t := "bcdf"
	// maxCost := int(3)

	// result: 1
	// s := "abcd"
	// t := "cdef"
	// maxCost := int(3)

	// result: 1
	s := "abcd"
	t := "acde"
	maxCost := int(0)

	// result: 
	// s := ""
	// t := ""
	// maxCost := int(0)

	result := equalSubstring(s, t, maxCost)
	fmt.Printf("result = %v\n", result)
}

