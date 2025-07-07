package main

import (
	heap "container/heap"
	"fmt"
	"sort"
)

type Heap []int

func (h Heap) Len() int {
	return len(h)
}

func (h Heap) Less(i, j int) bool {
	return h[i] < h[j]
}

func (h Heap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *Heap) Push(x any) {
	*h = append(*h, x.(int))
}

func (h *Heap) Pop() any {
	old := *h
	n := len(old)
	x := old[n - 1]
	*h = old[0 : n - 1]
	return x
}

func maxEvents(events [][]int) int {
	n := len(events)
	maxDay := int(0)

	for _, event := range events {
		maxDay = max(maxDay, event[1])
	}

	sort.Slice(events, func(i, j int) bool {
		return events[i][0] < events[j][0]
	})

	pq := &Heap{}
	heap.Init(pq)
	result := int(0)

	for i, j := 1, 0; i <= maxDay; i++ {
		for j < n && events[j][0] <= i {
			heap.Push(pq, events[j][1])
			j++
		}

		for pq.Len() > 0 && (*pq)[0] < i {
			heap.Pop(pq)
		}

		if pq.Len() > 0 {
			heap.Pop(pq)
			result++
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
	// result: 3
	// events := [][]int{{1,2},{2,3},{3,4}}

	// result: 4
	events := [][]int{{1,2},{2,3},{3,4},{1,2}}

	// result: 
	// events := [][]int{}

	result := maxEvents(events)
	fmt.Printf("result = %v\n", result)
}

