package main

import (
	"fmt"
)

func minDeletionSize(strs []string) int {
	if len(strs) == 0 {
		return 0
	}

	k := len(strs[0])
	result := int(0)

	for col := range k {
		for row := 1; row < len(strs); row++ {
			if strs[row][col] < strs[row - 1][col] {
				result++
				break
			}
		}
	}

	return result
}

func main() {
	// result: 1
	// strs := []string{"cba","daf","ghi"}

	// result: 0
	// strs := []string{"a","b"}

	// result: 3
	strs := []string{"zyx","wvu","tsr"}

	// result: 
	// strs := []string{}

	result := minDeletionSize(strs)
	fmt.Printf("result = %v\n", result)
}

