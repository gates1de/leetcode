package main
import (
	"fmt"
	"strings"
)

func numMagicSquaresInside(grid [][]int) int {
	result := int(0)
	m := len(grid)
	n := len(grid[0])

	for row := 0; row + 2 < m; row++ {
		for col := 0; col + 2 < n; col++ {
			if isMagicSquare(grid, row, col) {
				result++
			}
		}
	}

	return result
}

func isMagicSquare(grid [][]int, row int, col int) bool {
	seq := "2943816729438167";
	seqReversed := "7618349276183492"

	border := ""
	for _, num := range []int{0,1,2,5,8,7,6,3} {
		num := grid[row + num / 3][col + (num % 3)]
		border += fmt.Sprintf("%v", num)
	}

	return grid[row][col] % 2 == 0 &&
		(strings.Contains(seq, border) || strings.Contains(seqReversed, border)) &&
		grid[row + 1][col + 1] == 5
}

func main() {
	// result: 1
	// grid := [][]int{{4,3,8,4},{9,5,1,9},{2,7,6,2}}

	// result: 0
	grid := [][]int{{8}}

	// result: 
	// grid := [][]int{}

	result := numMagicSquaresInside(grid)
	fmt.Printf("result = %v\n", result)
}

