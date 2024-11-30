package main
import (
	"fmt"
)

func slidingPuzzle(board [][]int) int {
	target := [2][3]byte{{1, 2, 3}, {4, 5, 0}}
	initial := [2][3]byte{}
	for i, row := range board {
		for j, val := range row {
			initial[i][j] = byte(val)
		}
	}

	queue := []*[2][3]byte{&initial}
	visited := map[[2][3]byte]bool{initial: true}
	steps := int(0)

	for len(queue) > 0 {
		for _, board := range queue {
			queue = queue[1:]

			if *board == target {
				return steps
			}

			for r := 0; r < 2; r++ {
				for c := 0; c < 3; c++ {
					if board[r][c] != 0 {
						continue
					}

					if next := getNext(r, c, r + 1, c, board); next != nil && !visited[*next] {
						visited[*next] = true; queue = append(queue, next)
					}
					if next := getNext(r, c, r - 1, c, board); next != nil && !visited[*next] {
						visited[*next] = true; queue = append(queue, next)
					}

					if next := getNext(r, c, r, c + 1, board); next != nil && !visited[*next] {
						visited[*next] = true; queue = append(queue, next)
					}

					if next := getNext(r, c, r, c - 1, board); next != nil && !visited[*next] {
						visited[*next] = true; queue = append(queue, next) }
					}
				}
			}
			steps++
	}

	return -1
}

func getNext(sr, sc, dr, dc int, board *[2][3]byte) *[2][3]byte {
	if dr < 0 || dr >= 2 || dc < 0 || dc >= 3 {
		return nil
	}

	next := [2][3]byte{}

	for r := 0; r < 2; r++ {
		for c := 0; c < 3; c++ {
			next[r][c] = board[r][c]
		}
	}

	next[sr][sc], next[dr][dc] = next[dr][dc], next[sr][sc]
	return &next
}

func main() {
	// result: 1
	// board := [][]int{{1,2,3},{4,0,5}}

	// result: -1
	// board := [][]int{{1,2,3},{5,4,0}}

	// result: 5
	board := [][]int{{4,1,2},{5,0,3}}

	// result: 
	// board := [][]int{}

	result := slidingPuzzle(board)
	fmt.Printf("result = %v\n", result)
}

