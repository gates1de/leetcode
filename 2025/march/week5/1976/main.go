package main
import (
	"fmt"
	"math"
	heap "container/heap"
)

const modulo = int(1e9 + 7)

type Heap [][]int64

func (h Heap) Len() int           { return len(h) }
func (h Heap) Less(i, j int) bool { return h[i][0] < h[j][0] }
func (h Heap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h Heap) Peek() []int64      { return h[0] }

func (h *Heap) Push(x interface{}) {
	*h = append(*h, x.([]int64))
}

func (h *Heap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n - 1]
	*h = old[0 : n - 1]
	return x
}

func countPaths(n int, roads [][]int) int {
	graph := make([][][]int, n)
	for i, _ := range graph {
		graph[i] = make([][]int, 0)
	}

	for _, road := range roads {
		start := road[0]
		end := road[1]
		travelTime := road[2]
		graph[start] = append(graph[start], []int{end, travelTime})
		graph[end] = append(graph[end], []int{start, travelTime})
	}

	pq := &Heap{}
	heap.Init(pq)
	shortestTime := make([]int64, n)
	for i, _ := range shortestTime {
		shortestTime[i] = math.MaxInt64
	}

	pathCount := make([]int, n)
	shortestTime[0] = 0
	pathCount[0] = 1

	heap.Push(pq, []int64{0, 0})

	for pq.Len() > 0 {
		top := heap.Pop(pq).([]int64)
		currentTime := top[0]
		currentNode := int(top[1])

		if currentTime > shortestTime[currentNode] {
			continue
		}

		for _, neighbor := range graph[currentNode] {
			neighborNode := neighbor[0]
			roadTime := int64(neighbor[1])

			if currentTime + roadTime < shortestTime[neighborNode] {
				shortestTime[neighborNode] = currentTime + roadTime
				pathCount[neighborNode] = pathCount[currentNode]
				heap.Push(pq, []int64{shortestTime[neighborNode], int64(neighborNode)})
			} else if currentTime + roadTime == shortestTime[neighborNode] {
				pathCount[neighborNode] = (pathCount[neighborNode] + pathCount[currentNode]) % modulo
			}
		}
	}

	return pathCount[n - 1]
}

func main() {
	// result: 4
	// n := int(7)
	// roads := [][]int{{0,6,7},{0,1,2},{1,2,3},{1,3,3},{6,3,3},{3,5,1},{6,5,1},{2,5,1},{0,4,5},{4,6,2}}

	// result: 1
	n := int(2)
	roads := [][]int{{1,0,10}}

	// result: 
	// n := int(0)
	// roads := [][]int{}

	result := countPaths(n, roads)
	fmt.Printf("result = %v\n", result)
}

