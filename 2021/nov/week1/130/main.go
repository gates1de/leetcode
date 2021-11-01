package main
import (
	"fmt"
)

func solve(board [][]byte)  {
	m := len(board)
	n := len(board[0])

	if m == 0 {
		return
	}

	for i := 0; i < m; i++ {
		dfs(i, 0, board)
		dfs(i, n - 1, board)
	}

	for j := 0; j < n; j++ {
		dfs(0, j, board)
		dfs(m - 1, j, board)
	}

	for i := range board {
		for j := range board[i] {
			if board[i][j] == 'o' {
				board[i][j] = 'O'
			} else {
				board[i][j] = 'X'
			}
		}
	}
}

func dfs(i int, j int, board [][]byte) {
	if i < 0 || j < 0 || i >= len(board) || j >= len(board[0]) {
		return
	}

	if board[i][j] == 'O' {
		board[i][j] = 'o'
		dfs(i - 1, j, board)
		dfs(i + 1, j, board)
		dfs(i, j - 1, board)
		dfs(i, j + 1, board)
	}
}

// Out of memory
func ngSolution(board [][]byte)  {
	m := len(board)
	n := len(board[0])

	if m <= 2 || n <= 2 {
		return
	}

	oMap := map[[2]int]bool{}
	// up side & down side
	for x := 1; x < n - 1; x++ {
		// up to down
		// fmt.Printf("%v: up to down\n", string(rune(board[0][x])))
		if board[0][x] == 'O' {
			oCells := [][2]int{}
			for y := 1; y < m - 1; y++ {
				// fmt.Printf("board[%v][%v] = %v\n", y, x, board[y][x])
				if board[y][x] == 'X' {
					break
				}
				oCells = append(oCells, [2]int{y, x})
			}

			updateO(oCells, oMap, board)
		}

		// donw to up
		if board[m - 1][x] == 'O' {
			oCells := [][2]int{}
			for y := m - 2; y > 0; y-- {
				if board[y][x] == 'X' {
					break
				}
				oCells = append(oCells, [2]int{y, x})
			}

			updateO(oCells, oMap, board)
		}
	}

	// left side & right side
	for y := 1; y < m - 1; y++ {
		// left to right
		if board[y][0] == 'O' {
			oCells := [][2]int{}
			for x := 1; x < n - 1; x++ {
				if board[y][x] == 'X' {
					break
				}
				oCells = append(oCells, [2]int{y, x})
			}

			updateO(oCells, oMap, board)
		}

		// right to left
		if board[y][n - 1] == 'O' {
			oCells := [][2]int{}
			for x := n - 2; x > 0; x-- {
				if board[y][x] == 'X' {
					break
				}
				oCells = append(oCells, [2]int{y, x})
			}

			updateO(oCells, oMap, board)
		}
	}

	// fmt.Printf("oMap = %v\n", oMap)
	for i := 1; i < m - 1; i++ {
		for j := 1; j < n - 1; j++ {
			if !oMap[[2]int{i, j}] {
				board[i][j] = 'X'
			}
		}
	}
}

func updateO(cells [][2]int, oMap map[[2]int]bool, board [][]byte) {
	m := len(board)
	n := len(board[0])

	fmt.Printf("oCells = %v\n", cells)
	for len(cells) > 0 {
		cell := cells[0]
		cells = cells[1:]
		oMap[cell] = true
		y := cell[0]
		x := cell[1]
		// up
		if y - 1 > 0 && board[y - 1][x] == 'O' && !oMap[[2]int{y - 1, x}] {
			cells = append(cells, [2]int{y - 1, x})
		}

		// right
		if x + 1 < n - 1 && board[y][x + 1] == 'O' && !oMap[[2]int{y, x + 1}] {
			cells = append(cells, [2]int{y, x + 1})
		}

		// down
		if y + 1 < m - 1 && board[y + 1][x] == 'O' && !oMap[[2]int{y + 1, x}] {
			cells = append(cells, [2]int{y + 1, x})
		}

		// left
		if x - 1 > 0 && board[y][x - 1] == 'O' && !oMap[[2]int{y, x - 1}] {
			cells = append(cells, [2]int{y, x - 1})
		}
	}
}

func main() {
	// result: [["X","X","X","X"],["X","X","X","X"],["X","X","X","X"],["X","O","X","X"]]
	// board := [][]byte{
	// 	{'X','X','X','X'},
	// 	{'X','O','O','X'},
	// 	{'X','X','O','X'},
	// 	{'X','O','X','X'},
	// }

	// result: ["X"]
	// board := [][]byte{{'X'}}

	// result: [["X","X","X","X"],["X","O","O","X"],["X","O","O","X"],["X","O","X","X"]]
	// board := [][]byte{
	// 	{'X','X','X','X'},
	// 	{'X','O','O','X'},
	// 	{'X','O','O','X'},
	// 	{'X','O','X','X'},
	// }

	// result: [["X","O","X"],["X","O","O"]]
	// board := [][]byte{{'X','O','X'},{'X','O','O'}}

	// result: [["X","X","X","O","X"],["X","X","X","O","X"],["X","X","X","O","O"],["X","O","X","X","X"],["X","O","X","X","X"]]
	board := [][]byte{
		{'X','X','X','O','X'},
		{'X','O','X','O','X'},
		{'X','X','O','O','O'},
		{'X','O','X','X','X'},
		{'X','O','X','X','X'},
	}

	// result: 
	// board := [][]byte{}

	solve(board)
	fmt.Printf("result = %v\n", board)
}

