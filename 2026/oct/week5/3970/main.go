package main

import (
	"container/heap"
	"fmt"
)

func shortestPath(n int, edges [][]int, labels string, k int) int {
	if n == 0 {
		return -1
	}
	if n == 1 {
		return 0
	}

	type edge struct {
		to     int
		weight int64
	}

	graph := make([][]edge, n)
	for _, currentEdge := range edges {
		graph[currentEdge[0]] = append(graph[currentEdge[0]], edge{
			to:     currentEdge[1],
			weight: int64(currentEdge[2]),
		})
	}

	totalStates := n * k
	distances := make([]int64, totalStates)
	for i := range distances {
		distances[i] = -1
	}
	distances[0] = 0

	queue := &minHeap{{distance: 0, state: 0}}
	heap.Init(queue)

	for queue.Len() > 0 {
		current := heap.Pop(queue).(heapItem)
		if distances[current.state] != current.distance {
			continue
		}

		currentNode := current.state / k
		currentRun := current.state%k + 1
		if currentNode == n-1 {
			return int(current.distance)
		}

		for _, currentEdge := range graph[currentNode] {
			nextRun := 1
			if labels[currentNode] == labels[currentEdge.to] {
				nextRun = currentRun + 1
			}
			if nextRun > k {
				continue
			}

			nextState := currentEdge.to*k + nextRun - 1
			nextDistance := current.distance + currentEdge.weight
			if distances[nextState] == -1 || nextDistance < distances[nextState] {
				distances[nextState] = nextDistance
				heap.Push(queue, heapItem{
					distance: nextDistance,
					state:    nextState,
				})
			}
		}
	}

	return -1
}

type heapItem struct {
	distance int64
	state    int
}

type minHeap []heapItem

func (h minHeap) Len() int {
	return len(h)
}

func (h minHeap) Less(i int, j int) bool {
	return h[i].distance < h[j].distance
}

func (h minHeap) Swap(i int, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *minHeap) Push(value any) {
	*h = append(*h, value.(heapItem))
}

func (h *minHeap) Pop() any {
	old := *h
	last := len(old) - 1
	item := old[last]
	*h = old[:last]
	return item
}

func main() {
	// result: 3
	// n := 3
	// edges := [][]int{{0,1,1},{1,2,1},{0,2,3}}
	// labels := "aab"
	// k := 1

	// result: 2
	// n := 3
	// edges := [][]int{{0,1,1},{1,2,1},{0,2,3}}
	// labels := "aaa"
	// k := 2

	// result: -1
	n := 3
	edges := [][]int{{0, 1, 1}, {1, 2, 1}}
	labels := "aaa"
	k := 2

	// result: 0
	// n := 0
	// edges := [][]int{}
	// labels := ""
	// k := 0

	result := shortestPath(n, edges, labels, k)
	fmt.Printf("result = %v\n", result)
}
