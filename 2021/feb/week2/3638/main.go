package main
import (
	localHeap "./heap"
	"container/heap"
	"fmt"
)

func shortestPathBinaryMatrix(grid [][]int) int {
	if len(grid) <= 0 || len(grid[0]) <= 0 {
		return 0
	}
	if grid[0][0] == 1 {
		return -1
	}

	maxX := len(grid[0]) - 1
	maxY := len(grid) - 1
	h := &localHeap.Heap{}
	heap.Init(h)
	dir := [][]int{{0, 1}, {1, 1}, {1, 0}, {1, -1}, {0, -1}, {-1, -1}, {-1, 0}, {-1, 1}}

	visited := make([][]bool, len(grid))
	for i := range visited {
		visited[i] = make([]bool, len(grid[0]))
	}
	heap.Push(h, []int{0, 0, 0})

	for h.Len() > 0 {
		top := heap.Pop(h).([]int)
		if visited[top[0]][top[1]] {
			continue
		}
		visited[top[0]][top[1]] = true

		if top[0] == maxY && top[1] == maxX {
			return top[2] + 1
		}

		for _, d := range dir {
			newY, newX := top[0] + d[0], top[1] + d[1]

			if (newX < 0 || maxX < newX) || (newY < 0 || maxY < newY) || grid[newY][newX] == 1 {
				continue
			}

			heap.Push(h, []int{newY, newX, top[2] + 1})
		}
	}
	return -1
}

func max(a, b int) int {
	if b > a {
		return b
	}
	return a
}

func abs(a, b int) int {
	if b > a {
		return b - a
	}
	return a - b
}

func main() {
	// grid := [][]int{{0, 1}, {1, 0}}
	// grid := [][]int{{0, 0, 0}, {1, 1, 0}, {1, 1, 0}}
	grid := [][]int{{1, 0, 0}, {1, 1, 0}, {1, 1, 0}}
	result := shortestPathBinaryMatrix(grid)
	fmt.Printf("result = %v\n", result)
}

