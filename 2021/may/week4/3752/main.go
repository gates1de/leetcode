package main
import (
	"fmt"
)

const EMPTY = "-"
const FILLED = "."
const QUEEN = "Q"

func solveNQueens(n int) [][]string {
	result := [][]string{}
	for i := 0; i < n; i++ {
		board := generateBoard(n)
		board[0][i] = QUEEN
		fill(0, i, board)
		put(board, n - 1, &result)
	}
	return result
}

func put(board [][]string, numOfQueens int, result *[][]string) {
	// printBoard(board)
	// fmt.Printf("numOfQueens = %v\n", numOfQueens)
	if numOfQueens == 0 {
		// printBoard(board)
		rows := boardToRows(board)
		*result = append(*result, rows)
		return
	}

	TOP:
	for i, row := range board {
		for j, square := range row {
			// fmt.Printf("board[%v][%v] = %v\n", i, j, square)
			if square == QUEEN {
				continue TOP
			}
			if square == FILLED {
				continue
			}

			// can put
			copyboard := copyBoard(board)
			copyboard[i][j] = QUEEN
			fill(i, j, copyboard)
			// printBoard(copyboard)
			put(copyboard, numOfQueens - 1, result)
		}
		return
	}
}

func fill(i int, j int, board [][]string) {
	n := len(board)
	for ii, row2 := range board {
		// fill diagonal
		rowDiff := abs(i, ii)
		// [x, y]
		upleft := []int{j - rowDiff, i - rowDiff}
		upright := []int{j + rowDiff, i - rowDiff}
		downleft := []int{j - rowDiff, i + rowDiff}
		downright := []int{j + rowDiff, i + rowDiff}
		// fmt.Printf("upleft = %v, upright = %v, downleft = %v, downright = %v\n", upleft, upright, downleft, downright)
		if upleft[0] >= 0 && upleft[1] >= 0 && board[upleft[1]][upleft[0]] == EMPTY {
			board[upleft[1]][upleft[0]] = FILLED
		}
		if upright[0] < n && upright[1] >= 0 && board[upright[1]][upright[0]] == EMPTY {
			board[upright[1]][upright[0]] = FILLED
		}
		if downleft[0] >= 0 && downleft[1] < n && board[downleft[1]][downleft[0]] == EMPTY {
			board[downleft[1]][downleft[0]] = FILLED
		}
		if downright[0] < n && downright[1] < n && board[downright[1]][downright[0]] == EMPTY {
			board[downright[1]][downright[0]] = FILLED
		}
		// fill vertical
		if board[ii][j] == EMPTY {
			board[ii][j] = FILLED
		}
		// fill horizontal
		for jj, s := range row2 {
			if i == ii && s == EMPTY {
				board[ii][jj] = FILLED
			}
		}
	}
}

func generateBoard(n int) [][]string {
	board := make([][]string, n)
	for i, row := range board {
		row = make([]string, n)
		for j, _ := range row {
			row[j] = EMPTY
		}
		board[i] = row
	}
	return board
}

func copyBoard(board [][]string) [][]string {
	result := make([][]string, len(board))
	for i, row := range board {
		row = make([]string, len(board))
		for j, _ := range row {
			row[j] = board[i][j]
		}
		result[i] = row
	}
	return result
}

func boardToRows(board [][]string) []string {
	result := make([]string, len(board))
	for i, row := range board {
		rowString := ""
		for _, square := range row {
			rowString += square
		}
		result[i] = rowString
	}
	return result
}

func abs(a, b int) int {
	if b > a {
		return b - a
	}
	return a - b
}

func printBoard(board [][]string) {
	fmt.Printf("\n")
	for _, row := range board {
		fmt.Printf("%v\n", row)
	}
	fmt.Printf("\n")
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

