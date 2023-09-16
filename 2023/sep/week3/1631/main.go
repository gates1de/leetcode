package main
import (
	"fmt"
	"container/heap"
)

type Heap [][]int

func (h Heap) Len() int           { return len(h) }
func (h Heap) Less(a, b int) bool { return h[a][2] < h[b][2] }
func (h Heap) Swap(a, b int)      { h[a], h[b] = h[b], h[a] }
func (h Heap) Peek() []int        { return h[0] }

func (h *Heap) Push(x interface{}) {
    *h = append(*h, x.([]int))
}

func (h *Heap) Pop() interface{} {
    old := *h
    n := len(old)
    x := old[n-1]
    *h = old[0 : n-1]
    return x
}

func minimumEffortPath(heights [][]int) int {
    if len(heights) <= 0 || len(heights[0]) <= 0 {
        return 0
    }
    maxX := len(heights[0]) - 1
    maxY := len(heights) - 1
    h := &Heap{}
    heap.Init(h)
    dir := [][]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}}

    visited := make([][]bool, len(heights))
    for i := range visited {
        visited[i] = make([]bool, len(heights[0]))
    }
    heap.Push(h, []int{0, 0, 0})

    for h.Len() > 0 {
        top := heap.Pop(h).([]int)
        if visited[top[0]][top[1]] {
            continue
        }
        visited[top[0]][top[1]] = true

        if top[0] == maxY && top[1] == maxX {
            return top[2]
        }

        for _, d := range dir {
            newY, newX := top[0] + d[0], top[1] + d[1]

            if (newX < 0 || maxX < newX) || (newY < 0 || maxY < newY) {
                continue
            }

            heap.Push(h, []int{
                newY, newX, max(top[2],
                    abs(heights[newY][newX], heights[top[0]][top[1]])),
            })
        }
    }

    return 0
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
	// result: 2
	// heights := [][]int{{1,2,2},{3,8,2},{5,3,5}}

	// result: 1
	// heights := [][]int{{1,2,3},{3,8,4},{5,3,5}}

	// result: 0
	heights := [][]int{{1,2,1,1,1},{1,2,1,2,1},{1,2,1,2,1},{1,2,1,2,1},{1,1,1,2,1}}

	// result: 
	// heights := [][]int{}

	result := minimumEffortPath(heights)
	fmt.Printf("result = %v\n", result)
}

