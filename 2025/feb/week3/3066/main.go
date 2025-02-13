package main
import (
	"fmt"
	heap "container/heap"
)

type Heap []int

func (h Heap) Len() int           { return len(h) }
func (h Heap) Less(i, j int) bool { return h[i] < h[j] }
func (h Heap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h Heap) Peek() int          { return h[0] }

func (h *Heap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *Heap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func minOperations(nums []int, k int) int {
	pq := &Heap{}
	heap.Init(pq)

	for _, num := range nums {
		heap.Push(pq, num)
	}

	result := int(0)
	for pq.Peek() < k {
		x := heap.Pop(pq).(int)
		y := heap.Pop(pq).(int)
		heap.Push(pq, min(x, y) * 2 + max(x, y))
		result++
	}

	return result
}

func max(a, b int) int {
	if b > a {
		return b
	}
	return a
}

func min(a, b int) int {
	if b < a {
		return b
	}
	return a
}

func main() {
	// result: 2
	// nums := []int{2,11,10,1,3}
	// k := int(10)

	// result: 4
	nums := []int{1,1,2,4,9}
	k := int(20)

	// result: 
	// nums := []int{}
	// k := int(0)

	result := minOperations(nums, k)
	fmt.Printf("result = %v\n", result)
}

