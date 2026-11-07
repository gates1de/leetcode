package main

import (
	"container/heap"
	"fmt"
)

func assignTasks(servers []int, tasks []int) []int {
	result := make([]int, len(tasks))
	if len(tasks) == 0 {
		return result
	}

	free := make(serverHeap, len(servers))
	for index, weight := range servers {
		free[index] = server{weight: weight, index: index}
	}
	heap.Init(&free)

	busy := make(busyHeap, 0, len(servers))
	currentTime := int64(0)
	for taskIndex, duration := range tasks {
		currentTime = max(currentTime, int64(taskIndex))
		for busy.Len() > 0 && busy[0].freeAt <= currentTime {
			heap.Push(&free, server{weight: servers[busy[0].index], index: busy[0].index})
			heap.Pop(&busy)
		}

		if free.Len() == 0 {
			currentTime = busy[0].freeAt
			for busy.Len() > 0 && busy[0].freeAt <= currentTime {
				heap.Push(&free, server{weight: servers[busy[0].index], index: busy[0].index})
				heap.Pop(&busy)
			}
		}

		assigned := heap.Pop(&free).(server)
		result[taskIndex] = assigned.index
		heap.Push(&busy, busyServer{
			freeAt: currentTime + int64(duration),
			index:  assigned.index,
		})
	}

	return result
}

type server struct {
	weight int
	index  int
}

type serverHeap []server

func (h serverHeap) Len() int {
	return len(h)
}

func (h serverHeap) Less(i int, j int) bool {
	if h[i].weight == h[j].weight {
		return h[i].index < h[j].index
	}
	return h[i].weight < h[j].weight
}

func (h serverHeap) Swap(i int, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *serverHeap) Push(value any) {
	*h = append(*h, value.(server))
}

func (h *serverHeap) Pop() any {
	old := *h
	last := len(old) - 1
	value := old[last]
	*h = old[:last]
	return value
}

type busyServer struct {
	freeAt int64
	index  int
}

type busyHeap []busyServer

func (h busyHeap) Len() int {
	return len(h)
}

func (h busyHeap) Less(i int, j int) bool {
	return h[i].freeAt < h[j].freeAt
}

func (h busyHeap) Swap(i int, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *busyHeap) Push(value any) {
	*h = append(*h, value.(busyServer))
}

func (h *busyHeap) Pop() any {
	old := *h
	last := len(old) - 1
	value := old[last]
	*h = old[:last]
	return value
}

func main() {
	// result: [2,2,0,2,1,2]
	// servers := []int{3,3,2}
	// tasks := []int{1,2,3,2,1,2}

	// result: [1,4,1,4,1,3,2]
	servers := []int{5, 1, 4, 3, 2}
	tasks := []int{2, 1, 2, 4, 5, 2, 1}

	// result: []
	// servers := []int{}
	// tasks := []int{}

	result := assignTasks(servers, tasks)
	fmt.Printf("result = %v\n", result)
}
