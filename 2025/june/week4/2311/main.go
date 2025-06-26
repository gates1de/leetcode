package main

import (
	"fmt"
	"math/bits"
)

func longestSubsequence(s string, k int) int {
	sm := int(0)
	bits := bits.Len(uint(k))
	result := int(0)

	for i := 0; i < len(s); i++ {
		char := s[len(s)-1-i]

		if char == '1' {
			if i < bits && sm+(1<<i) <= k {
				sm += 1 << i
				result++
			}
		} else {
			result++
		}
	}

	return result
}

func main() {
	// result: 5
	// s := "1001010"
	// k := int(5)

	// result: 6
	s := "00101001"
	k := int(1)

	// result:
	// s := ""
	// k := int(0)

	result := longestSubsequence(s, k)
	fmt.Printf("result = %v\n", result)
}
