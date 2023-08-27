package main
import (
	"fmt"
)

func canCross(stones []int) bool {
	n := len(stones)
	mark := make(map[int]int)
	dp := make([][]bool, 2001)
	for i, _ := range dp {
		dp[i] = make([]bool, 2001)
	}
	for i, stone := range stones {
		mark[stone] = i
	}

	dp[0][0] = true
	for i, stone := range stones {
		for prevJump := 0; prevJump <= n; prevJump++ {
			if dp[i][prevJump] {
				if _, ok := mark[stone + prevJump]; ok {
					dp[mark[stone + prevJump]][prevJump] = true
				}
				if _, ok := mark[stone + prevJump + 1]; ok {
					dp[mark[stone + prevJump + 1]][prevJump + 1] = true
				}
				if _, ok := mark[stone + prevJump - 1]; ok {
					dp[mark[stone + prevJump - 1]][prevJump - 1] = true
				}
			}
		}
	}

	for i, _ := range stones {
		if dp[n - 1][i] {
			return true
		}
	}

	return false
}

func main() {
	// result: true
	// stones := []int{0,1,3,5,6,8,12,17}

	// result: false
	stones := []int{0,1,2,3,4,8,9,11}

	// result: 
	// stones := []int{}

	result := canCross(stones)
	fmt.Printf("result = %v\n", result)
}

