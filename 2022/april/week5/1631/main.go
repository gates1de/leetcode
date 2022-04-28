package main
import (
	"fmt"
	"math"
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

// Time Limit Exceeded
func ngSolution(heights [][]int) int {
	m := len(heights)
	n := len(heights[0])

	if m == 1 && n == 1 {
		return 0
	}

	memo := make([][]int, m)
	for i, _ := range memo {
		memo[i] = make([]int, n)
		for j, _ := range memo[i] {
			memo[i][j] = math.MaxInt32
		}
	}
	helper(0, 0, heights[0][0], 0, heights, memo)
	fmt.Println(memo)

	return memo[m - 1][n - 1]
}

func helper(x int, y int, pre int, min int, heights [][]int, memo [][]int) {
	fmt.Println(x, y, pre, min)
	m := len(heights)
	n := len(heights[0])

	current := heights[y][x]
	diff := abs(current, pre)
	if diff >= memo[y][x] || min >= memo[y][x] {
		return
	}
	if diff > min {
		min = diff
	}
	memo[y][x] = min

	if x == n - 1 && y == m - 1 {
		return
	}

	dirs := [][]int{{-1, 0}, {0, 1}, {1, 0}, {0, -1}}
	for _, dir := range dirs {
		nextY := y + dir[0]
		nextX := x + dir[1]
		if nextY >= 0 && nextY < m && nextX >= 0 && nextX < n {
			helper(nextX, nextY, current, min, heights, memo)
		}
	}
}

func abs(a, b int) int {
	if b > a {
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
	// heights := [][]int{
	// 	{1,2,1,1,1},
	// 	{1,2,1,2,1},
	// 	{1,2,1,2,1},
	// 	{1,2,1,2,1},
	// 	{1,1,1,2,1},
	// }

	// result: 0
	// heights := [][]int{{3}}

	// result: 1
	// heights := [][]int{{0},{1}}

	// result: 9
	heights := [][]int{{1,10,6,7,9,10,4,9}}

	// result: 
	// heights := [][]int{}

	result := minimumEffortPath(heights)
	fmt.Printf("result = %v\n", result)
}

