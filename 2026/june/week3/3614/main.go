package main

import (
	"fmt"
)

func processStr(s string, k int64) byte {
	n := len(s)
	lengths := make([]int64, n + 1)

	for i := range n {
		lengths[i + 1] = lengths[i]
		switch s[i] {
		case '*':
			if lengths[i + 1] > 0 {
				lengths[i + 1]--
			}
		case '#':
			lengths[i + 1] *= 2
		case '%':
		default:
			lengths[i + 1]++
		}
	}

	if k < 0 || k >= lengths[n] {
		return '.'
	}

	for i := n - 1; i >= 0; i-- {
		c := s[i]
		prevLen := lengths[i]
		curLen := lengths[i + 1]

		switch c {
		case '*':
			_ = curLen
		case '#':
			if k >= prevLen {
				k -= prevLen
			}
		case '%':
			k = curLen - 1 - k
		default:
			if k == curLen - 1 {
				return c
			}
		}
	}

	return '.'
}

func main() {
	// result: "a"
	// s := "a#b%*"
	// k := int64(1)

	// result: "d"
	// s := "cd%#*#"
	// k := int64(3)

	// result: "."
	s := "z*#"
	k := int64(0)

	// result: ""
	// s := ""
	// k := int64(0)

	result := processStr(s, k)
	fmt.Printf("result = %v\n", result)
}
