package main
import (
	"fmt"
)

const modulo = int(1e9 + 7)

func countRoutes(locations []int, start int, finish int, fuel int) int {
	n := len(locations)
	dp := make([][]int, n)
	for i, _ := range dp {
		dp[i] = make([]int, fuel + 1)
	}

	for i, _ := range dp[finish] {
		dp[finish][i] = 1
	}

	for i := 0; i <= fuel; i++ {
		for j := 0; j < n; j++ {
			for k := 0; k < n; k++ {
				if j == k {
					continue
				}
				if abs(locations[j], locations[k]) <= i {
					dp[j][i] = (dp[j][i] + dp[k][i - abs(locations[j], locations[k])]) % modulo
				}
			}
		}
	}

	return dp[start][fuel]
}

func abs(a, b int) int {
	if b > a {
		return b - a
	}
	return a - b
}

func main() {
	// result: 4
	// locations := []int{2,3,6,8,4}
	// start := int(1)
	// finish := int(3)
	// fuel := int(5)

	// result: 5
	// locations := []int{4,3,1}
	// start := int(1)
	// finish := int(0)
	// fuel := int(6)

	// result: 0
	locations := []int{5,2,1}
	start := int(0)
	finish := int(2)
	fuel := int(3)

	// result: 
	// locations := []int{}
	// start := int(0)
	// finish := int(0)
	// fuel := int(0)

	result := countRoutes(locations, start, finish, fuel)
	fmt.Printf("result = %v\n", result)
}

