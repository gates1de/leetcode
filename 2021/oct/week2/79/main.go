package main
import (
	"fmt"
)

func exist(board [][]byte, word string) bool {
    for x:= 0; x < len(board) ; x++ {
        for y := 0; y < len(board[0]); y++ {
            if board[x][y] == word[0] && dfs(board, word, x, y) {
                return true
            }
        }
    }
    return false
}

func dfs(board [][]byte, word string, x, y int) bool {
    if len(word) <= 1 {
        return true
    }

    tmp := board[x][y]
    board[x][y] = '0'
    if x - 1 >= 0 && board[x - 1][y] == word[1] && dfs(board, word[1:], x - 1, y) {
        return true
    }
    if y - 1 >= 0 && board[x][y - 1] == word[1] && dfs(board, word[1:], x, y - 1) {
        return true
    }
    if x + 1 < len(board) && board[x + 1][y] == word[1] && dfs(board, word[1:], x + 1, y) {
        return true
    }
    if y + 1 < len(board[0]) && board[x][y + 1] == word[1] && dfs(board, word[1:], x, y + 1) {
        return true
    }
    board[x][y] = tmp
    return false
}

// Slow & Use more memory
func mySolution(board [][]byte, word string) bool {
	for i, row := range board {
		for j, _ := range row {
			visited := make([][]bool, len(board))
			for i, _ := range board {
				visited[i] = make([]bool, len(board[i]))
			}
			if helper(i, j, visited, board, word) {
				return true
			}
		}
	}
	return false
}

func helper(y int, x int, visited [][]bool, board [][]byte, word string) bool {
	// fmt.Printf("word = %v\n", word)
	if len(word) == 0 || (len(word) == 1 && board[y][x] == word[0]) {
		return true
	}

	c := board[y][x]
	if c != word[0] {
		return false
	}

	visited[y][x] = true
	// fmt.Printf("[%v, %v]: visited = %v\n", y, x, visited)

    // up
    if y - 1 >= 0 && !visited[y - 1][x] {
		cv := copyVisited(visited)
		if helper(y - 1, x, cv, board, word[1:]) {
			return true
		}
    }
    // right
    if x + 1 < len(board[0]) && !visited[y][x + 1]{
		cv := copyVisited(visited)
		if helper(y, x + 1, cv, board, word[1:]) {
			return true
		}
    }
    // down
    if y + 1 < len(board) && !visited[y + 1][x] {
		cv := copyVisited(visited)
		if helper(y + 1, x, cv, board, word[1:]) {
			return true
		}
    }
    // left
    if x - 1 >= 0 && !visited[y][x - 1] {
		cv := copyVisited(visited)
		if helper(y, x - 1, cv, board, word[1:]) {
			return true
		}
    }

	return false
}

func copyVisited(visited [][]bool) [][]bool {
	copyVisited := make([][]bool, len(visited))
	for i, _ := range visited {
		copyVisited[i] = make([]bool, len(visited[i]))
		copy(copyVisited[i], visited[i])
	}
	return copyVisited
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

