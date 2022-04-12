package main
import (
	"fmt"
)

func gameOfLife(board [][]int)  {
	m := len(board)
	n := len(board[0])
	result := make([][]int, m)
	for i, _ := range result {
		result[i] = make([]int, n)
	}

	for i, row := range board {
		for j, own := range row {
			dirs := [][]int{
				{i - 1, j - 1},
				{i - 1, j},
				{i - 1, j + 1},
				{i, j - 1},
				{i, j + 1},
				{i + 1, j - 1},
				{i + 1, j},
				{i + 1, j + 1},
			}

			neighborCount := int(0)
			for _, dir := range dirs {
				if dir[0] < 0 || dir[0] >= m || dir[1] < 0 || dir[1] >= n {
					continue
				}
				
				if board[dir[0]][dir[1]] == 1 {
					neighborCount++
				}
			}

			result[i][j] = own
			if own == 1 {
				if neighborCount <= 1 || neighborCount >= 4 {
					result[i][j] = 0
				}
			} else {
				if neighborCount == 3 {
					result[i][j] = 1
				}
			}
		}
	}

	for i, row := range result {
		copy(board[i], row)
	}
}

func main() {
	// result: [[0,0,0],[1,0,1],[0,1,1],[0,1,0]]
	// board := [][]int{{0,1,0},{0,0,1},{1,1,1},{0,0,0}}

	// result: [[1,1],[1,1]]
	board := [][]int{{1,1},{1,0}}

	// result: 
	// board := [][]int{}

	gameOfLife(board)
	fmt.Printf("result = %v\n", board)
}

