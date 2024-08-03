package main
import (
	"fmt"
)

func minHeightShelves(books [][]int, shelfWidth int) int {
	n := len(books)
	dp := make([]int, n + 1)
	dp[0] = 0
	dp[1] = books[0][1]

	for i := 2; i <= n; i++ {
		remaining := shelfWidth - books[i - 1][0]
		maxHeight := books[i - 1][1]
		dp[i] = maxHeight + dp[i - 1]

		j := i - 1
		for j > 0 && remaining - books[j - 1][0] >= 0 {
			maxHeight = max(maxHeight, books[j - 1][1])
			remaining -= books[j - 1][0]
			dp[i] = min(dp[i], maxHeight + dp[j - 1])
			j--
		}
	}

	return dp[n]
}

func max(a, b int) int {
	if b > a {
		return b
	}
	return a
}

func min(a, b int) int {
	if b < a {
		return b
	}
	return a
}

func main() {
	// result: 6
	// books := [][]int{{1,1},{2,3},{2,3},{1,1},{1,1},{1,1},{1,2}}
	// shelfWidth := int(4)

	// result: 4
	books := [][]int{{1,3},{2,4},{3,2}}
	shelfWidth := int(6)

	// result: 
	// books := [][]int{}
	// shelfWidth := int(0)

	result := minHeightShelves(books, shelfWidth)
	fmt.Printf("result = %v\n", result)
}

