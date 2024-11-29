package main
import (
	"fmt"
	heap "container/heap"
)

type Heap [][]int

func (h Heap) Len() int { return len(h) }

func (h Heap) Less(a, b int) bool {
	return h[a][0] < h[b][0]
}

func (h Heap) Swap(a, b int) { h[a], h[b] = h[b], h[a] }

func (h *Heap) Push(x interface{}) {
	*h = append(*h, x.([]int))
}

func (h *Heap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n - 1]
	*h = old[0 : n - 1]
	return x
}

func minimumTime(grid [][]int) int {
	if grid[0][1] > 1 && grid[1][0] > 1 {
		return -1
	}

	m := len(grid)
	n := len(grid[0])
	dirs := [][]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}
	visited := make([][]bool, m)
	for i, _ := range visited {
		visited[i] = make([]bool, n)
	}

	priorityQueue := &Heap{}
	heap.Init(priorityQueue)
	heap.Push(priorityQueue, []int{grid[0][0], 0, 0})

	for priorityQueue.Len() > 0 {
		current := heap.Pop(priorityQueue).([]int)
		time := current[0]
		row := current[1]
		col := current[2]

		if row == m - 1 && col == n - 1 {
			return time
		}

		if visited[row][col] {
			continue
		}
		visited[row][col] = true

		for _, dir := range dirs {
			nextRow := row + dir[0]
			nextCol := col + dir[1]

			if !isValid(visited, nextRow, nextCol) {
				continue
			}

			waitTime := int(0)
			if (grid[nextRow][nextCol] - time) % 2 == 0 {
				waitTime = 1
			}

			nextTime := max(grid[nextRow][nextCol] + waitTime, time + 1)
			heap.Push(priorityQueue, []int{nextTime, nextRow, nextCol})
		}
	}

	return -1
}

func isValid(visited [][]bool, row int, col int) bool {
	return row >= 0 && col >= 0 && row < len(visited) && col < len(visited[0]) && !visited[row][col]
}

func max(a, b int) int {
	if b > a {
		return b
	}
	return a
}

func main() {
	// result: 7
	// grid := [][]int{{0,1,3,2},{5,1,2,5},{4,3,8,6}}

	// result: -1
	grid := [][]int{{0,2,4},{3,2,1},{1,0,4}}

	// result: 
	// grid := [][]int{}

	result := minimumTime(grid)
	fmt.Printf("result = %v\n", result)
}

