package main
import (
	localHeap "heap/heap"
	"container/heap"
	"fmt"
)

func longestIncreasingPath(matrix [][]int) int {
    if len(matrix) == 0 || len(matrix[0]) == 0 {
        return 0
    }

    m := len(matrix)
    n := len(matrix[0])

    memo := make([][]int, m)
    for i := 0; i < m; i++ {
        memo[i] = make([]int, n)
    }

    result := 1
    for i := 0; i < m; i++ {
        for j := 0; j < n; j++ {
            result = max(result, helper(matrix, i, j, memo))
        }
    }
    return result
}

func helper(matrix [][]int, i int, j int, memo [][]int) int {
    if memo[i][j] != 0 {
        return memo[i][j]
    }

    m := len(matrix)
    n := len(matrix[0])

    result := 1
	// left
    if i - 1 >= 0 && matrix[i - 1][j] > matrix[i][j] {
        result = max(result,  1 + helper(matrix, i - 1, j, memo))
    }

	// right
    if i + 1 < m && matrix[i + 1][j] > matrix[i][j] {
        result = max(result, 1 + helper(matrix, i + 1, j, memo))
    }

	// up
    if j - 1 >= 0 && matrix[i][j - 1] > matrix[i][j] {
        result = max(result, 1 + helper(matrix, i, j - 1, memo))
    }

	// down
    if j + 1 < n && matrix[i][j + 1] > matrix[i][j] {
        result = max(result, 1 + helper(matrix, i, j + 1, memo))
    }

    memo[i][j] = result
    return result
}

// Time Limit Exceeded
func ngSolution(matrix [][]int) int {
	if len(matrix) == 0 {
		return 0
	} else if len(matrix) == 1 && len(matrix[0]) == 1 {
		return 1
	}

	maxX := len(matrix[0]) - 1
	maxY := len(matrix) - 1
	h := &localHeap.Heap{}
	heap.Init(h)
	dir := [][]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}}

	visited := make([][]bool, len(matrix))
	for i := range visited {
		visited[i] = make([]bool, len(matrix[0]))
	}

	heap.Push(h, []int{0, 0, 0})
	count := int(0)
	for h.Len() > 0 {
		top := heap.Pop(h).([]int)
		// fmt.Printf("top = %v, current matrix = %v\n", top, matrix[top[0]][top[1]])

		visited[top[0]][top[1]] = true

		for _, d := range dir {
			newY, newX := top[0] + d[0], top[1] + d[1]
			if (newX < 0 || maxX < newX) || (newY < 0 || maxY < newY) || (matrix[newY][newX] <= matrix[top[0]][top[1]]) {
				continue
			}
			heap.Push(h, []int{newY, newX, top[2] + 1})
			if count < top[2] + 1 {
				count = top[2] + 1
			}
		}

		if h.Len() == 0 {
			TOP: for i, v := range visited {
				for j, isVisited := range v {
					if !isVisited {
						heap.Push(h, []int{i, j, 0})
						break TOP
					}
				}
			}
		}
	}
	return count + 1
}

func abs(a int, b int) int {
	if a - b < 0 {
		return b - a
	}
	return a - b
}

func max(a int, b int) int {
	if a < b {
		return b
	}
	return a
}

func main() {
	// result: 4 ( path -> [1, 2, 6, 9] )
	// matrix := [][]int{{9, 9, 4}, {6, 6, 8}, {2, 1, 1}}

	// result: 4 ( path -> [3, 4, 5, 6] )
	// matrix := [][]int{{3, 4, 5}, {3, 2, 6}, {2, 2, 1}}

	// result: 1
	// matrix := [][]int{{1}}

	// result: 2 ( path -> [1, 2] )
	matrix := [][]int{{1, 2}}

	result := longestIncreasingPath(matrix)
	fmt.Printf("result = %v\n", result)
}

