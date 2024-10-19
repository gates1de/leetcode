package main
import (
	"fmt"
	"math"
	heap "container/heap"
)

type Heap []int

func (h Heap) Len() int { return len(h) }

func (h Heap) Less(a, b int) bool {
	return h[a] > h[b]
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

func maxKelements(nums []int, k int) int64 {
	priorityQueue := &Heap{}
	heap.Init(priorityQueue)

	for _, num := range nums {
		heap.Push(priorityQueue, num)
	}

	result := int64(0)

	for i := 0; i < k; i++ {
		target := heap.Pop(priorityQueue).(int)

		result += int64(target)
		heap.Push(priorityQueue, int(math.Ceil(float64(target) / 3)))
	}

	return result
}

func main() {
	// result: 50
	// nums := []int{10,10,10,10,10}
	// k := int(5)

	// result: 17
	nums := []int{1,10,3,3,3}
	k := int(3)

	// result: 
	// nums := []int{}
	// k := int(0)

	result := maxKelements(nums, k)
	fmt.Printf("result = %v\n", result)
}

