package main

import (
	"fmt"
)

func stoneGameII(piles []int) int {
	n := len(piles)
	if n == 0 {
		return 0
	}

	suffix := make([]int, n+1)
	for i := n - 1; i >= 0; i-- {
		suffix[i] = suffix[i+1] + piles[i]
	}

	dp := make([][]int, n)
	for i := range dp {
		dp[i] = make([]int, n+1)
	}

	sameMQueue := make([]int, n)
	secondRangeQueue := make([]uint16, n*n)
	secondRangeHead := make([]int, n)
	secondRangeTail := make([]int, n)
	for m := n; m >= 1; m-- {
		sameHead, sameTail := 0, 0

		for i := n - 1; i >= 0; i-- {
			if next := i + 1; next < n {
				for sameHead < sameTail && dp[sameMQueue[sameTail-1]][m] >= dp[next][m] {
					sameTail--
				}
				sameMQueue[sameTail] = next
				sameTail++
			}
			for sameHead < sameTail && sameMQueue[sameHead] > i+m {
				sameHead++
			}

			head, tail := secondRangeHead[i], secondRangeTail[i]
			if x := m + 1; i+x < n {
				for head < tail {
					previousX := int(secondRangeQueue[i*n+tail-1])
					if dp[i+previousX][previousX] < dp[i+x][x] {
						break
					}
					tail--
				}
				secondRangeQueue[i*n+tail] = uint16(x)
				tail++
			}
			for head < tail && int(secondRangeQueue[i*n+head]) > 2*m {
				head++
			}
			secondRangeHead[i], secondRangeTail[i] = head, tail

			if i+2*m >= n {
				dp[i][m] = suffix[i]
				continue
			}

			minimum := dp[sameMQueue[sameHead]][m]
			candidateX := int(secondRangeQueue[i*n+head])
			if candidate := dp[i+candidateX][candidateX]; candidate < minimum {
				minimum = candidate
			}
			dp[i][m] = suffix[i] - minimum
		}
	}

	return dp[0][1]
}

func main() {
	// result: 10
	// piles := []int{2,7,9,4,4}

	// result: 104
	piles := []int{1, 2, 3, 4, 5, 100}

	// result:
	// piles := []int{}

	result := stoneGameII(piles)
	fmt.Printf("result = %v\n", result)
}
