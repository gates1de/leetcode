package main
import (
	"fmt"
	"math"
	heap "container/heap"
)

type Heap [][]int

func (h Heap) Len() int           { return len(h) }
func (h Heap) Less(i, j int) bool { return h[i][2] < h[j][2] }
func (h Heap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h Heap) Peek() []int        { return h[0] }

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

func maxPoints(grid [][]int, queries []int) []int {
	queryCount := len(queries)
	m := len(grid)
	n := len(grid[0])
	totalCells := m * n
	thresholdForMaxPoints := make([]int, totalCells + 1)
	minValueToReach := make([][]int, m)
	for i, _ := range minValueToReach {
		minValueToReach[i] = make([]int, n)
		for j, _ := range minValueToReach[i] {
			minValueToReach[i][j] = math.MaxInt32
		}
	}
	minValueToReach[0][0] = grid[0][0]

	pq := &Heap{}
	heap.Init(pq)
	heap.Push(pq, []int{0, 0, grid[0][0]})
	visitedCells := int(0)

	for pq.Len() > 0 {
		current := heap.Pop(pq).([]int)
		visitedCells++
		thresholdForMaxPoints[visitedCells] = current[2]

		for _, dir := range [][]int{{1,0},{-1,0},{0,1},{0,-1}} {
			newRow := current[0] + dir[0]
			newCol := current[1] + dir[1]

			if newRow >= 0 && newRow < m && newCol >= 0 && newCol < n && minValueToReach[newRow][newCol] == math.MaxInt32 {
				minValueToReach[newRow][newCol] = max(current[2], grid[newRow][newCol])
				heap.Push(pq, []int{newRow, newCol, minValueToReach[newRow][newCol]})
			}
		}
	}

	result := make([]int, queryCount)
	for i, query := range queries {
		left := int(0)
		right := totalCells

		for left < right {
			mid := (left + right + 1) >> 1
			if thresholdForMaxPoints[mid] < query {
				left = mid
			} else {
				right = mid - 1
			}
		}

		result[i] = left
	}

	return result
}

func max(a, b int) int {
	if b > a {
		return b
	}
	return a
}

func main() {
	// result: [5,8,1]
	// grid := [][]int{{1,2,3},{2,5,7},{3,5,1}}
	// queries := []int{5,6,2}

	// result: [0]
	grid := [][]int{{5,2,1},{1,1,2}}
	queries := []int{3}

	// result: []
	// grid := [][]int{}
	// queries := []int{}

	result := maxPoints(grid, queries)
	fmt.Printf("result = %v\n", result)
}

