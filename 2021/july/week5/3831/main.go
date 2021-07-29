package main
import (
	"fmt"
)

func updateMatrix(matrix [][]int) [][]int {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return matrix
	}

	n, m := len(matrix), len(matrix[0])
	coordinates := make([][2]int, 0)
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if matrix[i][j] == 0 {
				coordinates = append(coordinates, [2]int{i, j})
				continue
			}
			matrix[i][j] = 1 << 31 - 1
		}
	}

	ways := [][2]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}
	for len(coordinates) > 0 {
		current := coordinates[0]
		currentVal := matrix[current[0]][current[1]]
		coordinates = coordinates[1:]

		for _, way := range ways {
			i := current[0] + way[0]
			j := current[1] + way[1]

			if i < 0 || j < 0 || i == n || j == m || matrix[i][j] <= currentVal+1 {
				continue
			}
			matrix[i][j] = currentVal + 1
			coordinates = append(coordinates, [2]int{i, j})
		}
	}
	return matrix
}

// Time Limit Exceeded
func ngSolution(mat [][]int) [][]int {
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

