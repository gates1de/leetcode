package main
import (
	localHeap "heap/heap"
	"container/heap"
	"fmt"
)

func longestIncreasingPath(matrix [][]int) int {
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
		fmt.Printf("top = %v, current matrix = %v\n", top, matrix[top[0]][top[1]])
		// if visited[top[0]][top[1]] {
		// 	continue
		// }
		visited[top[0]][top[1]] = true
		// if top[0] == maxY && top[1] == maxX {
		// 	return top[2]
        // }
		for _, d := range dir {
			newY, newX := top[0] + d[0], top[1] + d[1]
			if (newX < 0 || maxX < newX) || (newY < 0 || maxY < newY) || (matrix[newY][newX] <= matrix[top[0]][top[1]]) {
				continue
			}
			heap.Push(h, []int{newY, newX, top[2] + 1})
			if count < top[2] + 1 {
				count = top[2] + 1
			}
			// heap.Push(h, []int{
			// 	newY, newX, max(
			// 		top[2],
			// 		abs(matrix[newY][newX], matrix[top[0]][top[1]]),
			// 	),
			// })
		}

		if h.Len() == 0 {
			fmt.Printf("heap = %v\n", h)
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
	fmt.Printf("visited = %v\n", visited)
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
	matrix := [][]int{{1}}

	// result: 2 ( path -> [1, 2] )
	// matrix := [][]int{{1, 2}}

	// result:
	// matrix := [][]int{}

	result := longestIncreasingPath(matrix)
	fmt.Printf("result = %v\n", result)
}

