package main

import (
	"container/heap"
	"fmt"
	"math"
)

type Step struct {
	ID   int
	Cost int
}

func NewStep(id int, costFromStart int) Step {
	return Step{ID: id, Cost: costFromStart}
}

type PriorityQueue []Step

func (pq PriorityQueue) Len() int {
    return len(pq)
}

func (pq PriorityQueue) Less(first int, second int) bool {
	return pq[first].Cost < pq[second].Cost
}

func (pq PriorityQueue) Swap(first int, second int) {
	pq[first], pq[second] = pq[second], pq[first]
}

func (pq *PriorityQueue) Push(object any) {
	step := object.(Step)
	*pq = append(*pq, step)
}

func (pq *PriorityQueue) Pop() any {
	step := (*pq)[pq.Len() - 1]
	*pq = (*pq)[0 : pq.Len() - 1]
	return step
}

func minCost(numberOfNodes int, edges [][]int) int {
	graph := createGraph(numberOfNodes, edges)
	start := int(0)
	goal := numberOfNodes - 1
	return findMinCostPath(graph, start, goal)
}

func findMinCostPath(graph []map[int]int, start int, goal int) int {
	pq := PriorityQueue{}
	minCostPath := make([]int, len(graph))
	for i := range minCostPath {
		minCostPath[i] = math.MaxInt
	}

	heap.Push(&pq, NewStep(start, 0))
	minCostPath[start] = 0
	for pq.Len() > 0 {
		current := heap.Pop(&pq).(Step)
		if current.ID == goal {
			return current.Cost
		}

		for next := range graph[current.ID] {
			if minCostPath[next] > current.Cost + graph[current.ID][next] {
				minCostPath[next] = current.Cost + graph[current.ID][next]
				heap.Push(&pq, NewStep(next, minCostPath[next]))
			}
		}
	}

	return -1
}

func createGraph(numberOfNodes int, edges [][]int) []map[int]int {
	graph := make([]map[int]int, numberOfNodes)
	for i := range numberOfNodes {
		graph[i] = map[int]int{}
	}

	for _, edge := range edges {
		from := edge[0]
		to := edge[1]
		cost := edge[2]

		if _, has := graph[from][to]; !has {
			graph[from][to] = cost
		}
		if _, has := graph[to][from]; !has {
			graph[to][from] = 2 * cost
		}

		graph[from][to] = min(graph[from][to], cost)
		graph[to][from] = min(graph[to][from], 2 * cost)
	}

	return graph
}

func main() {
	// result: 5
	// n := int(4)
	// edges := [][]int{{0,1,3},{3,1,1},{2,3,4},{0,2,2}}

	// result: 3
	n := int(4)
	edges := [][]int{{0,2,1},{2,1,1},{1,3,1},{2,3,3}}

	// result: 
	// n := int(0)
	// edges := [][]int{}

	result := minCost(n, edges)
	fmt.Printf("result = %v\n", result)
}

