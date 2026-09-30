package main

import (
	"fmt"
)

func maxDepthAfterSplit(seq string) []int {
	result := make([]int, len(seq))
	depth := 0
	for i := range seq {
		if seq[i] == '(' {
			depth++
			result[i] = depth & 1
		} else {
			result[i] = depth & 1
			depth--
		}
	}

	return result
}

func main() {
	// result: [0,1,1,1,1,0]
	// seq := "(()())"

	// result: [0,0,0,1,1,0,1,1]
	seq := "()(())()"

	// result: []
	// seq := ""

	result := maxDepthAfterSplit(seq)
	fmt.Printf("result = %v\n", result)
}
