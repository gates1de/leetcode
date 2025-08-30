package main

import (
	"fmt"
)

func isValidSudoku(board [][]byte) bool {
    for i, row := range board {
        m := map[byte]bool{}
        for j, b := range row {
            if b == byte('.') {
                continue
            }

            // horizontal check
            if m[b] {
                return false
            }
            m[b] = true

            // vertical check
            for ii, _ := range board {
                if ii == i {
                    continue
                }
                if board[ii][j] == b {
                    return false
                }
            }
        }
    }

    for i := 0; i < len(board); i += 3 {
        for j := 0; j < len(board[0]); j += 3 {
            if !threeMatCheck(i, j, board) {
                return false
            }
        }
    }
    return true
}

func threeMatCheck(i int, j int, board [][]byte) bool {
    m := map[byte]bool{}
    for ii := i; ii < i + 3; ii++ {
        for jj := j; jj < j + 3; jj++ {
            b := board[ii][jj]
            if b == byte('.') {
                continue
            }
            if m[b] {
                return false
            }
            m[b] = true
        }
    }
    return true
}

func main() {
	// result: true
	// board := [][]byte{
	// 	{'5','3','.','.','7','.','.','.','.'},
	// 	{'6','.','.','1','9','5','.','.','.'},
	// 	{'.','9','8','.','.','.','.','6','.'},
	// 	{'8','.','.','.','6','.','.','.','3'},
	// 	{'4','.','.','8','.','3','.','.','1'},
	// 	{'7','.','.','.','2','.','.','.','6'},
	// 	{'.','6','.','.','.','.','2','8','.'},
	// 	{'.','.','.','4','1','9','.','.','5'},
	// 	{'.','.','.','.','8','.','.','7','9'},
	// }

	// result: false
	board := [][]byte{
		{'8','3','.','.','7','.','.','.','.'},
		{'6','.','.','1','9','5','.','.','.'},
		{'.','9','8','.','.','.','.','6','.'},
		{'8','.','.','.','6','.','.','.','3'},
		{'4','.','.','8','.','3','.','.','1'},
		{'7','.','.','.','2','.','.','.','6'},
		{'.','6','.','.','.','.','2','8','.'},
		{'.','.','.','4','1','9','.','.','5'},
		{'.','.','.','.','8','.','.','7','9'},
	}

	// result: 
	// board := [][]byte{}

	result := isValidSudoku(board)
	fmt.Printf("result = %v\n", result)
}

