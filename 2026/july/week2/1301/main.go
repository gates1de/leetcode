package main

import (
	"fmt"
)

const modulo = int(1e9 + 7)

func pathsWithMaxScore(board []string) []int {
	n := len(board)
	if n == 0 {
		return []int{}
	}

	type state struct {
		score int
		count int
	}

	dp := make([][]state, n)
	for i := range n {
		dp[i] = make([]state, n)
		for j := range n {
			dp[i][j].score = -1
		}
	}

	dp[0][0] = state{score: 0, count: 1}

	for i := range n {
		for j := range n {
			if board[i][j] == 'X' || dp[i][j].count == 0 {
				continue
			}

			curScore := dp[i][j].score
			for _, next := range [][2]int{{i + 1, j}, {i, j + 1}, {i + 1, j + 1}} {
				x := next[0]
				y := next[1]
				if x >= n || y >= n || board[x][y] == 'X' {
					continue
				}

				add := int(0)
				if board[x][y] >= '1' && board[x][y] <= '9' {
					add = int(board[x][y] - '0')
				}
				nextScore := curScore + add
				if nextScore > dp[x][y].score {
					dp[x][y].score = nextScore
					dp[x][y].count = dp[i][j].count
				} else if nextScore == dp[x][y].score {
					dp[x][y].count += dp[i][j].count
					if dp[x][y].count >= modulo {
						dp[x][y].count -= modulo
					}
				}
			}
		}
	}

	result := dp[n - 1][n - 1]
	if result.count == 0 {
		return []int{0, 0}
	}

	return []int{result.score, result.count}
}

func main() {
	// result: [7,1]
	// board := []string{"E23","2X2","12S"}

	// result: [4,2]
	// board := []string{"E12","1X1","21S"}

	// result: [0,0]
	board := []string{"E11","XXX","11S"}

	// result: []
	// board := []string{}

	result := pathsWithMaxScore(board)
	fmt.Printf("result = %v\n", result)
}
