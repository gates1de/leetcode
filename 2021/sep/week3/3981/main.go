package main
import (
	"fmt"
)

const a = "A"
const b = "B"
const d = "Draw"
const p = "Pending"

func tictactoe(moves [][]int) string {
	board := make([][]string, 3)
	for i, _ := range board {
		board[i] = make([]string, 3)
		for j, _ := range board {
			board[i][j] = " "
		}
	}

	isATurn := true
	for _, move := range moves {
		if isATurn {
			board[move[0]][move[1]] = "X"
		} else {
			board[move[0]][move[1]] = "O"
		}
		isATurn = !isATurn
	}

	return judge(board)
}

func judge(board [][]string) string {
	// printBoard(board)
	vertical := [3][3]string{}

	result := d
	// horizontal check
	for i, row := range board {
		if row[0] == "X" && row[1] == "X" && row[2] == "X" {
			return a
		} else if row[0] == "O" && row[1] == "O" && row[2] == "O" {
			return b
		}

		vertical[0][i] = board[i][0]
		vertical[1][i] = board[i][1]
		vertical[2][i] = board[i][2]
	}

	// vertical check
	for _, v := range vertical {
		if v[0] == "X" && v[1] == "X" && v[2] == "X" {
			return a
		}
		if v[0] == "O" && v[1] == "O" && v[2] == "O" {
			return b
		}
		if v[0] == " " || v[1] == " " || v[2] == " " {
			result = p
		}
	}

	// diagonal check
	if (board[0][0] == "X" && board[1][1] == "X" && board[2][2] == "X") ||
		(board[0][2] == "X" && board[1][1] == "X" && board[2][0] == "X") {
		return a
	}
	if (board[0][0] == "O" && board[1][1] == "O" && board[2][2] == "O") ||
		(board[0][2] == "O" && board[1][1] == "O" && board[2][0] == "O") {
		return b
	}

	return result
}

func printBoard(board [][]string) {
	fmt.Printf("---------")
	for _, row := range board {
		fmt.Printf("\n")
		for _, square := range row {
			fmt.Printf(" %v ", square)
		}
	}
	fmt.Printf("\n---------\n")
}

func main() {
	// result: "A"
	// moves := [][]int{{0, 0}, {2, 0}, {1, 1}, {2, 1}, {2, 2}}

	// result: "B"
	// moves := [][]int{{0, 0}, {1, 1}, {0, 1}, {0, 2}, {1, 0}, {2, 0}}

	// result: "Draw"
	// moves := [][]int{{0, 0}, {1, 1}, {2, 0}, {1, 0}, {1, 2}, {2, 1}, {0, 1}, {0, 2}, {2, 2}}

	// result: "Pending"
	// moves := [][]int{{0, 0}, {1, 1}}

	// result: ""
	// moves := [][]int{}

	result := tictactoe(moves)
	fmt.Printf("result = %v\n", result)
}

