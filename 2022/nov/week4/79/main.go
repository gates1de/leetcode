package main
import (
	"fmt"
)

func exist(board [][]byte, word string) bool {
	m := len(board)
	if m == 0 {
		return false
	}
	n := len(board[0])

	visited := make([][]bool, m)
	for i, _ := range visited {
		visited[i] = make([]bool, n)
	}

	for i, _ := range board {
		for j, _ := range board[i] {
			if board[i][j] != word[0] {
				continue
			}

			if helper(i, j, visited, board, word[1:]) {
				return true
			}
		}
	}

	return false
}

func helper(y int, x int, visited [][]bool, board [][]byte, word string) bool {
	if len(word) == 0 {
		return true
	}

	m := len(board)
	if m == 0 {
		return false
	}
	n := len(board[0])

	newVisited := make([][]bool, m)
	for i, _ := range newVisited {
		newVisited[i] = make([]bool, n)
		copy(newVisited[i], visited[i])
	}
	newVisited[y][x] = true

	dirs := [][]int{{0, -1}, {1, 0}, {0, 1}, {-1, 0}}
	for _, dir := range dirs {
		newX := x + dir[0]
		newY := y + dir[1]

		if newX < 0 || n <= newX || newY < 0 || m <= newY || newVisited[newY][newX] || board[newY][newX] != word[0] {
			continue
		}

		if helper(newY, newX, newVisited, board, word[1:]) {
			return true
		}
	}

	return false
}

func main() {
	// result: true
	// board := [][]byte{
	// 	{'A', 'B', 'C', 'E'},
	// 	{'S', 'F', 'C', 'S'},
	// 	{'A', 'D', 'E', 'E'},
	// }
	// word := "ABCCED"

	// result: true
	// board := [][]byte{
	// 	{'A', 'B', 'C', 'E'},
	// 	{'S', 'F', 'C', 'S'},
	// 	{'A', 'D', 'E', 'E'},
	// }
	// word := "SEE"

	// result: false
	// board := [][]byte{
	// 	{'A', 'B', 'C', 'E'},
	// 	{'S', 'F', 'C', 'S'},
	// 	{'A', 'D', 'E', 'E'},
	// }
	// word := "ABCB"

	// result: true
	// board := [][]byte{
	// 	{'A', 'B', 'C', 'E'},
	// 	{'S', 'F', 'C', 'S'},
	// 	{'A', 'D', 'E', 'E'},
	// }
	// word := "ASFBCCEESE"

	// result: true
	board := [][]byte{
		{'A', 'B', 'C', 'E'},
		{'S', 'F', 'E', 'S'},
		{'A', 'D', 'E', 'E'},
	}
	word := "ABCEFSADEESE"

	// result: 
	// board := [][]byte{
	// 	{'', '', '', ''},
	// }
	// word := ""

	result := exist(board, word)
	fmt.Printf("result = %v\n", result)
}

