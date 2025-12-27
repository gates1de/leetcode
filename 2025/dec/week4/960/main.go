package main

import (
	"fmt"
)

func minDeletionSize(strs []string) int {
	w := len(strs[0])
	dp := make([]int, w)
	for i := range dp {
		dp[i] = 1
	}

	for i := w - 2; i >= 0; i-- {
		MID:
		for j := i + 1; j < w; j++ {
			for _, str := range strs {
				if str[i] > str[j] {
					continue MID
				}
			}

			dp[i] = max(dp[i], 1 + dp[j])
		}
	}

	kept := int(0)
	for _, val := range dp {
		kept = max(kept, val)
	}
	return w - kept
}

func main() {
	// result: 3
	// strs := []string{"babca","bbazb"}

	// result: 4
	// strs := []string{"edcba"}

	// result: 0
	strs := []string{"ghi","def","abc"}

	// result: 
	// strs := []string{}

	result := minDeletionSize(strs)
	fmt.Printf("result = %v\n", result)
}

