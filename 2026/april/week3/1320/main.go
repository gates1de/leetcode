package main

import (
	"fmt"
)

func minimumDistance(word string) int {
	n := len(word)
	dp := make([][]int, n)
	for i := range dp {
		dp[i] = make([]int, 26)
		for j := range dp[i] {
			dp[i][j] = 1 << 30
		}
	}

	for j := range 26 {
		dp[0][j] = 0
	}

	for i := 1; i < n; i++ {
		cur := int(word[i] - 'A')
		prev := int(word[i - 1] - 'A')
		d := getDistance(prev, cur)

		for j := range 26 {
			dp[i][j] = min(dp[i][j], dp[i - 1][j] + d)

			if prev == j {
				for k := range 26 {
					d0 := getDistance(k, cur)
					dp[i][j] = min(dp[i][j], dp[i - 1][k] + d0)
				}
			}
		}
	}

	result := 1 << 30
	for j := range 26 {
		result = min(result, dp[n - 1][j])
	}

	return result
}

func getDistance(p int, q int) int {
	x1, y1 := p / 6, p % 6
	x2, y2 := q / 6, q % 6
	return abs(x1 - x2) + abs(y1 - y2)
}

func abs(num int) int {
	if num < 0 {
		return -num
	}
	return num
}

func main() {
	// result: 3
	// word := "CAKE"

	// result: 6
	word := "HAPPY"

	// result: 
	// word := ""

	result := minimumDistance(word)
	fmt.Printf("result = %v\n", result)
}

