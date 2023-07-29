package main
import (
	"fmt"
	"math"
)

func soupServings(n int) float64 {
	m := int(math.Ceil(float64(n) / 25.0))
	dp := make(map[int]map[int]float64)
	for k := 1; k <= m; k++ {
		if calcDP(k, k, dp) > 1 - 1e-5 {
			return 1.0
		}
	}
	return calcDP(m, m, dp)
}

func calcDP(i int, j int, dp map[int]map[int]float64) float64 {
	if i <= 0 && j <= 0 {
		return 0.5
	}
	if i <= 0 {
		return 1.0
	}
	if j <= 0 {
		return 0.0
	}

	if dp[i] != nil {
		if _, ok := dp[i][j]; ok {
			return dp[i][j]
		}
	}

	result := calcDP(i - 4, j, dp) +
		calcDP(i - 3, j - 1, dp) +
		calcDP(i - 2, j - 2, dp) +
		calcDP(i - 1, j - 3, dp)
	result /= 4.0

	if dp[i] == nil {
		dp[i] = make(map[int]float64)
	}
	dp[i][j] = result
	return result
}

func main() {
	// result: 0.62500
	// n := int(50)

	// result: 0.71875
	n := int(100)

	// result: 
	// n := int(0)

	result := soupServings(n)
	fmt.Printf("result = %v\n", result)
}

