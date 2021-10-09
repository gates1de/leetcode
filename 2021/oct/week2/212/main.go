package main
import (
	"fmt"
)

func findWords(board [][]byte, words []string) []string {
	result := []string{}
	for x:= 0; x < len(board); x++ {
        for y := 0; y < len(board[0]); y++ {
			for i, word := range words {
				// fmt.Printf("[%v][%v] = %v: word = %v\n", x, y, board[x][y], word)
				if len(word) > 0 && board[x][y] == word[0] && dfs(board, word, x, y) {
					result = append(result, word)
					words[i] = ""
				}
			}
        }
    }

	return result
}

func dfs(board [][]byte, word string, x, y int) bool {
    if len(word) <= 1 {
        return true
    }

    tmp := board[x][y]
    board[x][y] = '0'
    if x - 1 >= 0 && board[x - 1][y] == word[1] && dfs(board, word[1:], x - 1, y) {
		board[x][y] = tmp
        return true
    }
    if y - 1 >= 0 && board[x][y - 1] == word[1] && dfs(board, word[1:], x, y - 1) {
		board[x][y] = tmp
        return true
    }
    if x + 1 < len(board) && board[x + 1][y] == word[1] && dfs(board, word[1:], x + 1, y) {
		board[x][y] = tmp
        return true
    }
    if y + 1 < len(board[0]) && board[x][y + 1] == word[1] && dfs(board, word[1:], x, y + 1) {
		board[x][y] = tmp
        return true
    }
	board[x][y] = tmp
    return false
}

func main() {
	// result: ["eat","oath"]
	// board := [][]byte{
	// 	{'o', 'a', 'a', 'n'},
	// 	{'e', 't', 'a', 'e'},
	// 	{'i', 'h', 'k', 'r'},
	// 	{'i', 'f', 'l', 'v'},
	// }
	// words := []string{"oath", "pea", "eat", "rain"}

	// result: ["abcdefg","befa","eaabcdgfa","gfedcbaaa"]
	board := [][]byte{
		{'a', 'b', 'c'},
		{'a', 'e', 'd'},
		{'a', 'f', 'g'},
	}
	words := []string{"abcdefg", "gfedcbaaa", "eaabcdgfa", "befa", "dgc", "ade"}

	// result: 
	// board := [][]byte{
	// 	{'', '', '', ''},
	// }
	// words := []string{}

	result := findWords(board, words)
	fmt.Printf("result = %v\n", result)
}

