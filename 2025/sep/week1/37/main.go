package main

import (
	"fmt"
)

const dot byte = '.'

func solveSudoku(board [][]byte)  {
    helper(0, 0, &board)
}

func helper(i, j int, board *[][]byte) bool {
    if j == 9 {
        j = 0
        i++
    }
    if i == 9 {
        return true
    }
    if (*board)[i][j] != dot {
        return helper(i, j + 1, board)
    }

    for candidate := byte('1'); candidate <= byte('9'); candidate++ {
        if isValid(i, j, candidate, *board) {
            (*board)[i][j] = candidate
            if helper(i, j + 1, board) {
                return true
            }
        }
    }
    (*board)[i][j] = '.'
    return false
}

func isValid(i, j int, input byte, board [][]byte) bool {
    // validate a row
    for pointer := 0; pointer < 9; pointer++ {
        if board[i][pointer] == input {
            return false
        }
    }

    // validate a col
    for pointer := 0; pointer < 9; pointer++ {
        if board[pointer][j] == input {
            return false
        }
    }


    // validate a square
    startRow, startCol := (i / 3) * 3, (j / 3) * 3;
    for row := startRow; row < startRow + 3; row++ {
        for col := startCol; col < startCol + 3; col++ {
            if board[row][col] == input {
                return false
            }
        }
    }
    return true

}

// Wrong Answer
func ngSolution(board [][]byte)  {
	m := map[[2]int]map[byte]bool{}
	remainCount := int(0)
	for i, row := range board {
		for j, b := range row {
			m[[2]int{i, j}] = map[byte]bool{}
			if b != byte('.') {
				remainCount++
				continue
			}

			setNums := map[byte]bool{}
			for ii, _ := range board {
				if ii == i || board[ii][j] == byte('.') {
					continue
				}
				setNums[board[ii][j]] = true
			}
			for jj, _ := range row {
				if jj == j || board[i][jj] == byte('.') {
					continue
				}
				setNums[board[i][jj]] = true
			}
			m[[2]int{i, j}] = setNums
		}
	}

	for i := 0; i < len(board); i += 3 {
		for j := 0; j < len(board[0]); j += 3 {
			threeMatCheck(i, j, m, board)
		}
	}

	TOP:
	for true {
		max := int(0)
		maxCoord := [2]int{-1, -1}
		for coord, setNums := range m {
			if board[coord[0]][coord[1]] != byte('.') || len(setNums) == 9 {
				continue
			}
			if len(setNums) == 8 {
				put(coord[0], coord[1], m, board)
				continue TOP
			}
			if max < len(setNums) {
				max = len(setNums)
				maxCoord = coord
			}
		}
		if max > 0 {
			put(maxCoord[0], maxCoord[1], m, board)
			continue
		}
		break
	}
}

func threeMatCheck(i int, j int, m map[[2]int]map[byte]bool, board [][]byte) {
	m[[2]int{i, j}] = map[byte]bool{}
	nums := map[byte]bool{}
	for ii := i; ii < i + 3; ii++ {
		for jj := j; jj < j + 3; jj++ {
			b := board[ii][jj]
			if b == byte('.') {
				continue
			}
			nums[b] = true
		}
	}
	for b, _ := range nums {
		for ii := i; ii < i + 3; ii++ {
			for jj := j; jj < j + 3; jj++ {
				m[[2]int{ii, jj}][b] = true
			}
		}
	}
}

func put(i int, j int, m map[[2]int]map[byte]bool, board[][]byte) {
	setNums := m[[2]int{i, j}]
	// fmt.Printf("[%v, %v]: setNums = %v\n", i, j, setNums)
	target := byte(' ')
	for i := 49; i <= 57; i++ {
		if setNums[byte(i)] {
			continue
		}
		target = byte(i)
		break
	}

	board[i][j] = target
	m[[2]int{i, j}][target] = true
	for ii := 0; ii < 9; ii++ {
		m[[2]int{ii, j}][target] = true
	}
	for jj := 0; jj < 9; jj++ {
		m[[2]int{i, jj}][target] = true
	}
	matRowIndex := 3 * (i / 3)
	matColIndex := 3 * (j / 3)
	for ii := matRowIndex; ii < matRowIndex + 3; ii++ {
		for jj := matColIndex; jj < matColIndex + 3; jj++ {
			m[[2]int{ii, jj}][target] = true
		}
	}
}

func strMatToByteMat(mat [][]string) [][]byte {
	result := make([][]byte, len(mat))
	for i, _ := range mat {
		result[i] = make([]byte, len(mat[i]))
		for j, s := range mat[i] {
			result[i][j] = []byte(s)[0]
		}
	}
	return result
}

func printBoard(board [][]byte) {
	fmt.Printf("[\n")
	for _, row := range board {
		fmt.Printf("[")
		comma := ""
		for _, b := range row {
			if b == byte('.') {
				fmt.Printf("%v.", comma)
			} else {
				fmt.Printf("%v%v", comma, b - 48)
			}
			comma = ", "
		}
		fmt.Printf("]\n")
	}
	fmt.Printf("\n]")
}

func main() {
	// result: [["5","3","4","6","7","8","9","1","2"],["6","7","2","1","9","5","3","4","8"],["1","9","8","3","4","2","5","6","7"],["8","5","9","7","6","1","4","2","3"],["4","2","6","8","5","3","7","9","1"],["7","1","3","9","2","4","8","5","6"],["9","6","1","5","3","7","2","8","4"],["2","8","7","4","1","9","6","3","5"],["3","4","5","2","8","6","1","7","9"]]
	// boardStr := [][]string{
	// 	{"5","3",".",".","7",".",".",".","."},
	// 	{"6",".",".","1","9","5",".",".","."},
	// 	{".","9","8",".",".",".",".","6","."},
	// 	{"8",".",".",".","6",".",".",".","3"},
	// 	{"4",".",".","8",".","3",".",".","1"},
	// 	{"7",".",".",".","2",".",".",".","6"},
	// 	{".","6",".",".",".",".","2","8","."},
	// 	{".",".",".","4","1","9",".",".","5"},
	// 	{".",".",".",".","8",".",".","7","9"},
	// }

	// result: [["5","1","9","7","4","8","6","3","2"],["7","8","3","6","5","2","4","1","9"],["4","2","6","1","3","9","8","7","5"],["3","5","7","9","8","6","2","4","1"],["2","6","4","3","1","7","5","9","8"],["1","9","8","5","2","4","3","6","7"],["9","7","5","8","6","3","1","2","4"],["8","3","2","4","9","1","7","5","6"],["6","4","1","2","7","5","9","8","3"]]
	// boardStr := [][]string{
	// 	{".",".","9","7","4","8",".",".","."},
	// 	{"7",".",".",".",".",".",".",".","."},
	// 	{".","2",".","1",".","9",".",".","."},
	// 	{".",".","7",".",".",".","2","4","."},
	// 	{".","6","4",".","1",".","5","9","."},
	// 	{".","9","8",".",".",".","3",".","."},
	// 	{".",".",".","8",".","3",".","2","."},
	// 	{".",".",".",".",".",".",".",".","6"},
	// 	{".",".",".","2","7","5","9",".","."},
	// }

	// result: [["8","5","4","2","1","9","7","6","3"],["3","9","7","8","6","5","4","2","1"],["2","6","1","4","7","3","9","8","5"],["7","8","5","1","2","6","3","9","4"],["6","4","9","5","3","8","1","7","2"],["1","3","2","9","4","7","8","5","6"],["9","2","6","3","8","4","5","1","7"],["5","1","3","7","9","2","6","4","8"],["4","7","8","6","5","1","2","3","9"]]
	boardStr := [][]string{
		{".",".",".","2",".",".",".","6","3"},
		{"3",".",".",".",".","5","4",".","1"},
		{".",".","1",".",".","3","9","8","."},
		{".",".",".",".",".",".",".","9","."},
		{".",".",".","5","3","8",".",".","."},
		{".","3",".",".",".",".",".",".","."},
		{".","2","6","3",".",".","5",".","."},
		{"5",".","3","7",".",".",".",".","8"},
		{"4","7",".",".",".","1",".",".","."},
	}

	// result: 
	// boardStr := [][]string{
	// }

	board := strMatToByteMat(boardStr)
	solveSudoku(board)
	printBoard(board)
}

