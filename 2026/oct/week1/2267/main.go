package main

import (
	"fmt"
)

func hasValidPath(grid [][]byte) bool {
	m := len(grid)
	if m == 0 {
		return false
	}
	n := len(grid[0])
	pathLength := m + n - 1
	if pathLength%2 == 1 || grid[0][0] == ')' || grid[m-1][n-1] == '(' {
		return false
	}

	words := (pathLength + 64) / 64
	dp := make([][4]uint64, n)

	for row := range grid {
		for col := range grid[row] {
			var source [4]uint64
			if row == 0 && col == 0 {
				source[0] = 1
			} else {
				if row > 0 {
					source = dp[col]
				}
				if col > 0 {
					for word := range words {
						source[word] |= dp[col-1][word]
					}
				}
			}

			var current [4]uint64
			if grid[row][col] == '(' {
				for word := words - 1; word >= 0; word-- {
					current[word] = source[word] << 1
					if word > 0 {
						current[word] |= source[word-1] >> 63
					}
				}
			} else {
				for word := range words {
					current[word] = source[word] >> 1
					if word+1 < words {
						current[word] |= source[word+1] << 63
					}
				}
			}
			dp[col] = current
		}
	}

	return dp[n-1][0]&1 == 1
}

func main() {
	// result: true
	// grid := [][]byte{{'(','(','('},{')','(',')'},{'(','(',')'},{'(','(',')'}}

	// result: false
	grid := [][]byte{{')', ')'}, {'(', '('}}

	// result: false
	// grid := [][]byte{}

	result := hasValidPath(grid)
	fmt.Printf("result = %v\n", result)
}
