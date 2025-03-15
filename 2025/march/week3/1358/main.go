package main
import (
	"fmt"
)

func numberOfSubstrings(s string) int {
	n := len(s)
	lastPosition := []int{-1, -1, -1}
	result := int(0)

	for i := 0; i < n; i++ {
		lastPosition[s[i] - 'a'] = i
		result += 1 + min(lastPosition[0], min(lastPosition[1], lastPosition[2]))
	}

	return result
}

func min(a, b int) int {
	if b < a {
		return b
	}
	return a
}

func main() {
	// result: 10
	// s := "abcabc"

	// result: 3
	// s := "aaacb"

	// result: 1
	s := "abc"

	// result: 
	// s := ""

	result := numberOfSubstrings(s)
	fmt.Printf("result = %v\n", result)
}

