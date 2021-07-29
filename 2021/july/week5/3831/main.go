package main
import (
	"fmt"
)

func updateMatrix(mat [][]int) [][]int {
	result := make([][]int, len(mat))
	for i, _ := range result {
		result[i] = make([]int, len(mat[0]))
		for j, _ := range result[i] {
			result[i][j] = -1
		}
	}
	for i, row := range mat {
		for j, num := range row {
			if num == 0 {
				result[i][j] = 0
				continue
			}
			visited := initVisited(len(mat), len(mat[0]))
			result[i][j] = helper(i, j, 0, visited, result, mat)
		}
	}
	return result
}

func initVisited(n int, m int) [][]int {
	visited := make([][]int, n)
	for i, _ := range visited {
		visited[i] = make([]int, m)
		for j, _ := range visited[i] {
			visited[i][j] = 10000
		}
	}
	return visited
}

func helper(i int, j int, step int, visited [][]int, result [][]int, mat [][]int) int {
	if visited[i][j] < 10000 {
		if visited[i][j] > step {
			visited[i][j] = step
		} else {
			return 10000
		}
	}

	if step > 0 {
		visited[i][j] = step
	}

	if mat[i][j] == 0 {
		return step
	}

	if result[i][j] >= 0 {
		return step + result[i][j]
	}

	maxY := len(mat)
	maxX := len(mat[0])

    // up
	minStep := int(10000)
    if i - 1 >= 0 {
		s := helper(i - 1, j, step + 1, visited, result, mat)
		if minStep > s {
			minStep = s
		}
    }

    // right
    if j + 1 < maxX {
		s := helper(i, j + 1, step + 1, visited, result, mat)
		if minStep > s {
			minStep = s
		}
    }

    // down
    if i + 1 < maxY {
		s := helper(i + 1, j, step + 1, visited, result, mat)
		if minStep > s {
			minStep = s
		}
    }

    // left
    if j - 1 >= 0 {
        s := helper(i, j - 1, step + 1, visited, result, mat)
		if minStep > s {
			minStep = s
		}
    }

	return minStep
}

func main() {
	// result: [[0,0,0],[0,1,0],[0,0,0]]
	// mat := [][]int{{0,0,0},{0,1,0},{0,0,0}}

	// result: [[0,0,0],[0,1,0],[1,2,1]]
	// mat := [][]int{{0,0,0},{0,1,0},{1,1,1}}

	// result: 
	mat := [][]int{{0},{1}}

	// result: 
	// mat := [][]int{}

	result := updateMatrix(mat)
	fmt.Printf("result = %v\n", result)
}

