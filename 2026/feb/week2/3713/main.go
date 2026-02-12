package main

import (
	"fmt"
)

func longestBalanced(s string) int {
	n := len(s)
	result := int(0)

	for i := range n {
		counter := make([]int, 26)

		for j := i; j < n; j++ {
			c := s[j] - 'a'
			counter[c]++
			flag := true

			for _, x := range counter {
				if x > 0 && x != counter[c] {
					flag = false
					break
				}
			}

			if flag && (j - i + 1) > result {
				result = j - i + 1
			}
		}
	}

	return result
}

func main() {
	// result: 4
	// s := "abbac"

	// result: 4
	// s := "zzabccy"

	// result: 2
	s := "aba"

	// result: 
	// s := ""

	result := longestBalanced(s)
	fmt.Printf("result = %v\n", result)
}

