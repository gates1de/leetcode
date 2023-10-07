package main
import (
	"fmt"
)

const modulo = int64(1e9 + 7)

func numOfArrays(n int, m int, k int) int {
	dp := make([][][]int64, n + 1)
	prefix := make([][][]int64, n + 1)
	for i, _ := range dp {
		dp[i] = make([][]int64, m + 1)
		prefix[i] = make([][]int64, m + 1)
		for j, _ := range dp[i] {
			dp[i][j] = make([]int64, k + 1)
			prefix[i][j] = make([]int64, k + 1)
		}
	}

	for num := 1; num <= m; num++ {
		dp[1][num][1] = 1
		prefix[1][num][1] = prefix[1][num - 1][1] + 1
	}

	for i := 1; i <= n; i++ {
		for maxNum := 1; maxNum <= m; maxNum++ {
			for cost := 1; cost <= k; cost++ {
				result := (int64(maxNum) * dp[i - 1][maxNum][cost]) % modulo
				result = (result + prefix[i - 1][maxNum - 1][cost - 1]) % modulo

				dp[i][maxNum][cost] += result
				dp[i][maxNum][cost] %= modulo

				prefix[i][maxNum][cost] = prefix[i][maxNum - 1][cost] + dp[i][maxNum][cost]
				prefix[i][maxNum][cost] %= modulo
			}
		}
	}

	return int(prefix[n][m][k])
}

func main() {
	// result: 6
	// n := int(2)
	// m := int(3)
	// k := int(1)

	// result: 0
	// n := int(5)
	// m := int(2)
	// k := int(3)

	// result: 1
	n := int(9)
	m := int(1)
	k := int(1)

	// result: 
	// n := int(0)
	// m := int(0)
	// k := int(0)

	result := numOfArrays(n, m, k)
	fmt.Printf("result = %v\n", result)
}

