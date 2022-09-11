package main
import (
	"fmt"
	"container/heap"
	"sort"

	localHeap "heap/heap"
)

const modulo = int(1e9 + 7)

func maxPerformance(n int, speed []int, efficiency []int, k int) int {
	sets := make([][]int, n)
	for i, _ := range sets {
		sets[i] = []int{efficiency[i], speed[i]}
	}

	sort.Slice(sets, func(a, b int) bool {
		return sets[a][0] > sets[b][0]
	})

	minHeap := &localHeap.MinHeap{}
	heap.Init(minHeap)
	result := int(0)
	sumSpeed := int(0)
	for _, set := range sets {
		eff := set[0]
		spd := set[1]
		heap.Push(minHeap, spd)
		sumSpeed += spd 
		if minHeap.Len() > k {
			sumSpeed -= heap.Pop(minHeap).(int)
		}
		result = max(result, sumSpeed * eff)
	}

	return result % modulo
}

func max(a, b int) int {
	if b > a {
		return b
	}
	return a
}

func main() {
	// result: 60
	n := int(6)
	speed := []int{2,10,3,1,5,8}
	efficiency := []int{5,4,3,9,7,2}
	k := int(2)

	// result: 
	// n := int(0)
	// speed := []int{}
	// efficiency := []int{}
	// k := int(0)

	result := maxPerformance(n, speed, efficiency, k)
	fmt.Printf("result = %v\n", result)
}

