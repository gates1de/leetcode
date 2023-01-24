package main
import (
	"fmt"
)

func snakesAndLadders(board [][]int) int {
	n := len(board)
	cells := make([][]int, n * n + 1)
	label := int(1)
	columns := make([]int, n)
	for i := 0; i < n; i++ {
		columns[i] = i
	}
	for row := n - 1; row >= 0; row-- {
		for _, column := range columns {
			cells[label] = []int{row, column}
			label++
		}
		reverse(columns)
	}

	dist := make([]int, n * n + 1)
	for i, _ := range dist {
		dist[i] = -1
	}
	queue := make([]int, 0, n * n + 1)
	dist[1] = 0
	queue = append(queue, 1)
	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]

		for next := current + 1; next <= min(current + 6, n * n); next++ {
			row := cells[next][0]
			column := cells[next][1]
			destination := next
			if board[row][column] != -1 {
				destination = board[row][column]
			}
			if dist[destination] == -1 {
				dist[destination] = dist[current] + 1
				queue = append(queue, destination)
			}
		}
	}

	return dist[n * n]
}

func min(a, b int) int {
	if b < a {
		return b
	}
	return a
}

func reverse(nums []int) {
	for i := 0; i < len(nums) / 2; i++ {
		nums[i], nums[len(nums) - i - 1] = nums[len(nums) - i - 1], nums[i]
	}
}

func main() {
	// result: 4
	// board := [][]int{
	// 	{-1,-1,-1,-1,-1,-1},
	// 	{-1,-1,-1,-1,-1,-1},
	// 	{-1,-1,-1,-1,-1,-1},
	// 	{-1,35,-1,-1,13,-1},
	// 	{-1,-1,-1,-1,-1,-1},
	// 	{-1,15,-1,-1,-1,-1},
	// }

	// result: 1
	board := [][]int{{-1,-1},{-1,3}}

	// result: 
	// board := [][]int{}

	result := snakesAndLadders(board)
	fmt.Printf("result = %v\n", result)
}

