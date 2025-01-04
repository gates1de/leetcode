package main
import (
	"fmt"
)

func countPalindromicSubsequence(s string) int {
	first := make([]int, 26)
	last := make([]int, 26)
	for i, _ := range first {
		first[i] = -1
	}

	for i, char := range s {
		curr := char - 'a'
		if first[curr] == -1 {
			first[curr] = i
		}

		last[curr] = i
	}

	result := int(0)

	for i, f := range first {
		if f == -1 {
			continue
		}

		l := last[i]
		between := make(map[byte]bool)
		for j := f + 1; j < l; j++ {
			between[s[j]] = true
		}

		result += len(between)
	}

	return result
}

func main() {
	// result: 3
	// s := "aabca"

	// result: 0
	// s := "adc"

	// result: 4
	s := "bbcbaba"

	// result: 
	// s := ""

	result := countPalindromicSubsequence(s)
	fmt.Printf("result = %v\n", result)
}
