package main
import (
	"fmt"
)

func rotateTheBox(box [][]byte) [][]byte {
	m := len(box)
	n := len(box[0])
	result := make([][]byte, n)

	for i, _ := range result {
		result[i] = make([]byte, m)

		for j, _ := range result[i] {
			result[i][j] = box[j][i]
		}
	}

	for i, _ := range result {
		for j := 0; j < m / 2; j++ {
			result[i][j], result[i][m - 1 - j] = result[i][m - 1 - j], result[i][j]
		}
	}

	for j := 0; j < m; j++ {
		lowestRowWithEmptyCell := n - 1

		for i := n - 1; i >= 0; i-- {
			if result[i][j] == '#' {
				result[i][j] = '.'
				result[lowestRowWithEmptyCell][j] = '#'
				lowestRowWithEmptyCell--
			}

			if result[i][j] == '*' {
				lowestRowWithEmptyCell = i - 1
			}
		}
	}

	return result
}

func main() {
	// result: [['.'],['#'],['#']]
	// box := [][]byte{{'#','.','#'}}

	// result: [['#','.'],['#','#'],['*','*'],['.','.']]
	// box := [][]byte{{'#','.','*','.'},{'#','#','*','.'}}

	// result: [['.','#','#'],['.','#','#'],['#','#','*'],['#','*','.'],['#','.','*'],['#','.','.']]
	box := [][]byte{{'#','#','*','.','*','.'},{'#','#','#','*','.','.'},{'#','#','#','.','#','.'}}

	// result: []
	// box := [][]byte{}

	result := rotateTheBox(box)
	fmt.Printf("result = %v\n", result)
}

