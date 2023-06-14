package main
import (
	"fmt"
	"math"
	"sort"
)

func minCost(n int, cuts []int) int {
	m := len(cuts)
	newCuts := make([]int, m + 2)
	copy(newCuts, cuts)
	newCuts[m + 1] = n
	sort.Ints(newCuts)
        // System.arraycopy(cuts, 0, newCuts, 1, m);
	dp := make([][]int, m + 2)
	for i, _ := range dp {
		dp[i] = make([]int, m + 2)
	}

	for diff := 2; diff < m + 2; diff++ {
		for left := 0; left < m + 2 - diff; left++ {
			right := left + diff
			val := math.MaxInt32
			for mid := left + 1; mid < right; mid++ {
				val = min(val, dp[left][mid] + dp[mid][right] + newCuts[right] - newCuts[left])
			}
			dp[left][right] = val
		}
	}

	return dp[0][m + 1]
}

func min(a, b int) int {
	if b < a {
		return b
	}
	return a
}

func main() {
	// result: 16
	// n := int(7)
	// cuts := []int{1,3,4,5}

	// result: 22
	n := int(9)
	cuts := []int{5,6,1,4,2}

	// result: 
	// n := int(0)
	// cuts := []int{}

	result := minCost(n, cuts)
	fmt.Printf("result = %v\n", result)
}

