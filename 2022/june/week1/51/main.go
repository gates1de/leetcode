package main
import (
	"fmt"
	"strings"
)

const Q = byte('Q')
const dot = byte('.')

func solveNQueens(n int) [][]string {
	result := [][]string{}
	board := make([]string, n)
	for i, _ := range board {
		board[i] = strings.Repeat(string(dot), n)
	}
	helper(0, n, board, n, &result)
	return result
}

func helper(row int, remain int, board []string, n int, result *[][]string) {
	if remain == 0 {
		*result = append(*result, board)
		return
	}

	TOP:
	for i := 0; i < n; i++ {
		for _, col := range board {
			if col[i] == Q {
				continue TOP
			}
		}

		// up left
		j := row - 1 
		k := i - 1
		for j >= 0 && k >= 0 {
			if board[j][k] == Q {
				continue TOP
			}
			j--
			k--
		}

		// down left
		j = row + 1
		k = i - 1
		for j < n && k >= 0 {
			if board[j][k] == Q {
				continue TOP
			}
			j++
			k--
		}

		// up right
		j = row - 1
		k = i + 1
		for j >= 0 && k < n {
			if board[j][k] == Q {
				continue TOP
			}
			j--
			k++
		}

		// down right
		j = row + 1
		k = i + 1
		for j < n && k < n {
			if board[j][k] == Q {
				continue TOP
			}
			j++
			k++
		}

		newBoard := make([]string, n)
		copy(newBoard, board)
		newBoard[row] = board[row][:i] + string(Q) + board[row][i + 1:]
		helper(row + 1, remain - 1, newBoard, n, result)
	}
}

func main() {
	// result: [[".Q..","...Q","Q...","..Q."],["..Q.","Q...","...Q",".Q.."]]
	// n := int(4)

	// result: [["Q"]]
	// n := int(1)

	// result: [["Q....","..Q..","....Q",".Q...","...Q."],["Q....","...Q.",".Q...","....Q","..Q.."],[".Q...","...Q.","Q....","..Q..","....Q"],[".Q...","....Q","..Q..","Q....","...Q."],["..Q..","Q....","...Q.",".Q...","....Q"],["..Q..","....Q",".Q...","...Q.","Q...."],["...Q.","Q....","..Q..","....Q",".Q..."],["...Q.",".Q...","....Q","..Q..","Q...."],["....Q",".Q...","...Q.","Q....","..Q.."],["....Q","..Q..","Q....","...Q.",".Q..."]]
	n := int(5)

	// result: 
	// n := int(0)

	result := solveNQueens(n)
	fmt.Printf("result = %v\n", result)
}

