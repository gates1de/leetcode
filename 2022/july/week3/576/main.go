package main
import (
	"fmt"
)

const modulo = int(1e9 + 7)

func findPaths(m int, n int, maxMove int, startRow int, startColumn int) int {
    dp := make([][][]int, maxMove + 1)
    for i := range dp {
        dp[i] = make([][]int, m)
        for j := range dp[i] {
            dp[i][j] = make([]int, n)
        }
    }

    dirs := [][]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}

    for s := 1; s <= maxMove; s++ {
        for y := 0; y < m; y++ {
            for x := 0; x < n; x++ {
                for _, dir := range dirs {
                    ty := dir[0] + y
                    tx := dir[1] + x
                    if tx < 0 || ty < 0 || tx >= n || ty >= m {
                        dp[s][y][x] += 1
                    } else {
                        dp[s][y][x] = (dp[s][y][x] + dp[s - 1][ty][tx]) % modulo
                    }
                }
            }
        }
    }

    return dp[maxMove][startRow][startColumn]
}

// Time Limit Exceeded
func ngSolution(m int, n int, maxMove int, startRow int, startColumn int) int {
	result := int(0)
	helper(startRow, startColumn, m, n, maxMove, &result)
	return result % modulo
}

func helper(y int, x int, m int, n int, remainingMove int, result *int) {
	if remainingMove == -1 {
		return
	}

	if y < 0 || m <= y || x < 0 || n <= x {
		*result++
		return
	}

	dirs := [][]int{{-1, 0}, {0, 1}, {1, 0}, {0, -1}}
	for _, dir := range dirs {
		helper(y + dir[0], x + dir[1], m, n, remainingMove - 1, result)
	}
}

func main() {
	// result: 6
	// m := int(2)
	// n := int(2)
	// maxMove := int(2)
	// startRow := int(0)
	// startColumn := int(0)

	// result: 12
	// m := int(1)
	// n := int(3)
	// maxMove := int(3)
	// startRow := int(0)
	// startColumn := int(1)

	// result: 102984580
	m := int(8)
	n := int(7)
	maxMove := int(16)
	startRow := int(1)
	startColumn := int(5)

	// result: 
	// m := int(0)
	// n := int(0)
	// maxMove := int(0)
	// startRow := int(0)
	// startColumn := int(0)

	result := findPaths(m, n, maxMove, startRow, startColumn)
	fmt.Printf("result = %v\n", result)
}

