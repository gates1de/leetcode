package main
import (
	"fmt"
	heap "container/heap"
	"math"
	"sort"
)

type Heap []int

func (h Heap) Len() int { return len(h) }

func (h Heap) Less(a, b int) bool {
	return h[a]< h[b]
}

func (h Heap) Swap(a, b int) { h[a], h[b] = h[b], h[a] }

func (h *Heap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *Heap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n - 1]
	*h = old[0 : n - 1]
	return x
}

func pickGifts(gifts []int, k int) int64 {
	n := len(gifts)
	sortedGifts := make([]int, n)
	copy(sortedGifts, gifts)
	sort.Ints(sortedGifts)

	priorityQueue := &Heap{}
	heap.Init(priorityQueue)

	for _, gift := range gifts {
		heap.Push(priorityQueue, -gift)
	}

	for i := 0; i < k; i++ {
		maxNum := -heap.Pop(priorityQueue).(int)
		heap.Push(priorityQueue, -int(math.Sqrt(float64(maxNum))))
	}

	result := int64(0)
	for priorityQueue.Len() > 0 {
		result -= int64(heap.Pop(priorityQueue).(int))
	}

	return result
}

func main() {
	// result: 29
	// gifts := []int{25,64,9,4,100}
	// k := int(4)

	// result: 4
	gifts := []int{1,1,1,1}
	k := int(4)

	// result: 
	// gifts := []int{}
	// k := int(0)

	result := pickGifts(gifts, k)
	fmt.Printf("result = %v\n", result)
}

