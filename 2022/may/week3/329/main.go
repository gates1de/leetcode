package main
import (
	"fmt"
)

func longestIncreasingPath(matrix [][]int) int {
	n := len(matrix)
	m := len(matrix[0])

	visited := make([][]int, n)
	for i, _ := range visited {
		visited[i] = make([]int, m)
	}

	result := int(0)
	for i, row := range matrix {
		for j, _ := range row {
			helper(j, i, 1, matrix, visited, &result)
		}
	}
	return result
}

func helper(x int, y int, count int, matrix [][]int, visited [][]int, result *int) {
	if visited[y][x] > 0 && visited[y][x] >= count {
		return
	}

	if *result < count {
		*result = count
	}
	visited[y][x] = count

	currentVal := matrix[y][x]
	n := len(matrix)
	m := len(matrix[0])

	// right, down, left, up
	dirs := [][]int{{1, 0}, {0, 1}, {-1, 0}, {0, -1}}

	for _, dir := range dirs {
		nextX := x + dir[0]
		nextY := y + dir[1]
		if nextX < 0 || m <= nextX || nextY < 0 || n <= nextY || currentVal >= matrix[nextY][nextX] {
			continue
		}

		helper(nextX, nextY, count + 1, matrix, visited, result)
	}
}

func main() {
	// result: 4
	// matrix := [][]int{{9,9,4},{6,6,8},{2,1,1}}

	// result: 4
	// matrix := [][]int{{3,4,5},{3,2,6},{2,2,1}}

	// result: 1
	// matrix := [][]int{{1}}

	// result: 2
	matrix := [][]int{{1, 2}}

	// result: 
	// matrix := [][]int{}

	result := longestIncreasingPath(matrix)
	fmt.Printf("result = %v\n", result)
}

