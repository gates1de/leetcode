package main
import (
	"fmt"
)

func movesToChessboard(board [][]int) int {
	isOK := check(board)
	if isOK {
		return 0
	}

	result := int(1000)
	changed := map[[3]int]bool{}
	helper(0, changed, board, &result)
	if result == 1000 {
		return -1
	}
	return result
}

func helper(count int, changed map[[3]int]bool, board [][]int, result *int) {
	fmt.Printf("board = %v\n", board)
	if *result <= count {
		return
	}
	if check(board) {
		if *result > count {
			*result = count
		}
		return
	}

	n := len(board)
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			if board[i][0] != board[j][0] && !changed[[3]int{0, i, j}] {
				newBoard := copyBoard(board)
				changed[[3]int{0, i, j}] = true
				horizontalSwap(i, j, newBoard)
				// fmt.Printf("newBoard = %v\n", newBoard)
				helper(count + 1, changed, newBoard, result)
			}
			if board[0][i] != board[0][j] && !changed[[3]int{1, i, j}] {
				newBoard := copyBoard(board)
				changed[[3]int{1, i, j}] = true
				verticalSwap(i, j, newBoard)
				// fmt.Printf("newBoard = %v\n", newBoard)
				helper(count + 1, changed, newBoard, result)
			}
		}
	}
}

func check(board [][]int) bool {
	expected := int(0)
	if board[0][0] == 1 {
		expected = 1
	}

	n := len(board)
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if board[i][j] != expected {
				return false
			}

			if expected == 0 {
				expected = 1
			} else {
				expected = 0
			}
		}

		if expected == 0 {
			expected = 1
		} else {
			expected = 0
		}
	}

	return true
}

func horizontalSwap(row1 int, row2 int, board [][]int) {
	n := len(board)
	for i := 0; i < n; i++ {
		board[row1][i], board[row2][i] = board[row2][i], board[row1][i]
	}
}

func verticalSwap(col1 int, col2 int, board [][]int) {
	n := len(board)
	for i := 0; i < n; i++ {
		board[i][col1], board[i][col2] = board[i][col2], board[i][col1]
	}
}

func copyBoard(board [][]int) [][]int {
	result := make([][]int, len(board))
	for i, row := range board {
		result[i] = make([]int, len(row))
		copy(result[i], row)
	}
	return result
}

func main() {
	// result: 2
	// board := [][]int{
	// 	{0,1,1,0},
	// 	{0,1,1,0},
	// 	{1,0,0,1},
	// 	{1,0,0,1},
	// }

	// result: 0
	// board := [][]int{{0, 1}, {1, 0}}

	// result: -1
	// board := [][]int{{1, 0}, {1, 0}}

	// result: 0
	board := [][]int{
		{0,1,0,1},
		{1,0,1,0},
		{0,1,0,1},
		{1,0,1,0},
	}

	// result: 
	// board := [][]int{}

	result := movesToChessboard(board)
	fmt.Printf("result = %v\n", result)
}

