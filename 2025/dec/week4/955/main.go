package main

import (
	"fmt"
)

func minDeletionSize(strs []string) int {
	n := len(strs)
	w := len(strs[0])
	cuts := make([]bool, n - 1)
	result := int(0)

	TOP:
	for j := range w {
		for i := range n - 1 {
			if !cuts[i] && strs[i][j] > strs[i + 1][j] {
				result++
				continue TOP
			}
		}

		for i := range n - 1 {
			if strs[i][j] < strs[i + 1][j] {
				cuts[i] = true
			}
		}
	}

	return result
}

func main() {
	// result: 1
	// strs := []string{"ca","bb","ac"}

	// result: 0
	// strs := []string{"xc","yb","za"}

	// result: 3
	strs := []string{"zyx","wvu","tsr"}

	// result: 
	// strs := []string{}

	result := minDeletionSize(strs)
	fmt.Printf("result = %v\n", result)
}

