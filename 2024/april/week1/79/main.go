package main
import (
	"fmt"
)

func exist(board [][]byte, word string) bool {
	m := len(board)
	n := len(board[0])
	visited := make([][]bool, m)
	for i, _ := range visited {
		visited[i] = make([]bool, n)
	}

	result := false

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if board[i][j] == word[0] {
				result = backtrack(i, j, 0, board, word, visited)

				if result {
					return true
				}
			}
		}
	}

	return false
}

func backtrack(i int, j int, index int, board [][]byte, word string, visited [][]bool) bool {
	m := len(board)
	n := len(board[0])
	if index == len(word) {
		return true
	}

	if i < 0 || i >= m || j < 0 || j >= n || visited[i][j] || board[i][j] != word[index] {
		return false
	}

	visited[i][j] = true

	if backtrack(i + 1, j, index + 1, board, word, visited) ||
		backtrack(i - 1, j, index + 1, board, word, visited) ||
		backtrack(i, j + 1, index + 1, board, word, visited) ||
		backtrack(i, j - 1, index + 1, board, word, visited) {
		return true
	}

	visited[i][j] = false
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

