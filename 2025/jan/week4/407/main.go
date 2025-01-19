package main
import (
	"fmt"
	heap "container/heap"
	"math"
)

type Grid struct {
	X int
	Y int
	Height int
}

type MinHeap []Grid

func (h MinHeap) Len() int {
	return len(h)
}

func (h MinHeap) Less(i int, j int) bool {
	return h[i].Height < h[j].Height
}

func (h MinHeap) Swap(i int, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *MinHeap) Push(a interface{}) {
	*h = append(*h, a.(Grid))
}

func (h *MinHeap) Pop() interface{} {
	l := len(*h)
	result := (*h)[l - 1]
	*h = (*h)[:l - 1]
	return result
}

func trapRainWater(heightMap [][]int) int {
	offset := [4][2]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}
	a := heightMap
	if len(a) == 0 || len(a[0]) == 0 {
		return 0
	}

	m := len(a)
	n := len(a[0])

	visited := make([][]bool, m)
	for i := 0; i < m; i++ {
		visited[i] = make([]bool, n)
	}

	h := &MinHeap{}
	for i := 0; i < m; i++ {
		heap.Push(h, Grid{i, 0, a[i][0]})
		visited[i][0] = true
		heap.Push(h, Grid{i, n - 1, a[i][n - 1]})
		visited[i][n - 1] = true
	}

	for i := 1; i < n - 1; i++ {
		heap.Push(h, Grid{0, i, a[0][i]})
		visited[0][i] = true
		heap.Push(h, Grid{m - 1, i, a[m - 1][i]})
		visited[m - 1][i] = true
	}

	result := int(0)
	waterLevel := math.MinInt64
	for h.Len() != 0 {
		low := heap.Pop(h).(Grid)
		waterLevel = max(waterLevel, a[low.X][low.Y])

		for _, v := range offset {
			g := Grid{low.X + v[0], low.Y + v[1], 0}

			if g.X < 0 || g.X >= m ||
				g.Y < 0 || g.Y >= n || visited[g.X][g.Y] {
				continue
			}

			g.Height = a[g.X][g.Y]

			if a[g.X][g.Y] < waterLevel {
				result += waterLevel - g.Height
			}

			visited[g.X][g.Y] = true
			heap.Push(h, g)
		}
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
	// result: 4
	// heightMap := [][]int{{1,4,3,1,3,2},{3,2,1,3,2,4},{2,3,3,2,3,1}}

	// result: 10
	heightMap := [][]int{{3,3,3,3,3},{3,2,2,2,3},{3,2,1,2,3},{3,2,2,2,3},{3,3,3,3,3}}

	// result: 
	// heightMap := [][]int{}

	result := trapRainWater(heightMap)
	fmt.Printf("result = %v\n", result)
}

