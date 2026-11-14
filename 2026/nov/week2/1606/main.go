package main

import (
	"container/heap"
	"fmt"
)

func busiestServers(k int, arrival []int, load []int) []int {
	if k == 0 {
		return []int{}
	}

	available := newServerTree(k)
	busy := make(busyHeap, 0, k)
	requests := make([]int, k)

	for i, time := range arrival {
		currentTime := int64(time)
		for busy.Len() > 0 && busy[0].finish <= currentTime {
			server := heap.Pop(&busy).(busyServer)
			available.update(server.id, server.id)
		}

		preferred := i % k
		server := available.query(preferred, k-1)
		if server == -1 {
			server = available.query(0, preferred-1)
		}
		if server == -1 {
			continue
		}

		requests[server]++
		available.update(server, available.inf)
		heap.Push(&busy, busyServer{
			finish: currentTime + int64(load[i]),
			id:     server,
		})
	}

	maxRequests := 0
	for _, requestCount := range requests {
		maxRequests = max(maxRequests, requestCount)
	}

	result := make([]int, 0)
	for server, requestCount := range requests {
		if requestCount == maxRequests {
			result = append(result, server)
		}
	}

	return result
}

type busyServer struct {
	finish int64
	id     int
}

type busyHeap []busyServer

func (h busyHeap) Len() int {
	return len(h)
}

func (h busyHeap) Less(i int, j int) bool {
	return h[i].finish < h[j].finish
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

type serverTree struct {
	tree []int
	size int
	inf  int
}

func newServerTree(k int) serverTree {
	size := 1
	for size < k {
		size <<= 1
	}

	inf := k
	tree := make([]int, size*2)
	for i := range tree {
		tree[i] = inf
	}
	for i := 0; i < k; i++ {
		tree[size+i] = i
	}
	for i := size - 1; i > 0; i-- {
		tree[i] = min(tree[i*2], tree[i*2+1])
	}

	return serverTree{tree: tree, size: size, inf: inf}
}

func (s *serverTree) update(index int, value int) {
	position := s.size + index
	s.tree[position] = value
	for position >>= 1; position > 0; position >>= 1 {
		s.tree[position] = min(s.tree[position*2], s.tree[position*2+1])
	}
}

func (s *serverTree) query(left int, right int) int {
	if left > right {
		return -1
	}

	left += s.size
	right += s.size + 1
	result := s.inf
	for left < right {
		if left&1 == 1 {
			result = min(result, s.tree[left])
			left++
		}
		if right&1 == 1 {
			right--
			result = min(result, s.tree[right])
		}
		left >>= 1
		right >>= 1
	}

	if result == s.inf {
		return -1
	}
	return result
}

func main() {
	// result: [1]
	// k := 3
	// arrival := []int{1,2,3,4,5}
	// load := []int{5,2,3,3,3}

	// result: [0]
	// k := 3
	// arrival := []int{1,2,3,4}
	// load := []int{1,2,1,2}

	// result: [0,1,2]
	k := 3
	arrival := []int{1, 2, 3}
	load := []int{10, 12, 11}

	// result: []
	// k := 0
	// arrival := []int{}
	// load := []int{}

	result := busiestServers(k, arrival, load)
	fmt.Printf("result = %v\n", result)
}
