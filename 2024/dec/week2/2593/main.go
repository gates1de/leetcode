package main
import (
	"fmt"
	heap "container/heap"
)

type Heap [][]int

func (h Heap) Len() int { return len(h) }

func (h Heap) Less(a, b int) bool {
	if h[a][0] != h[b][0] {
		return h[a][0] < h[b][0]
	}
	return h[a][1] < h[b][1]
}

func (h Heap) Swap(a, b int) { h[a], h[b] = h[b], h[a] }

func (h *Heap) Push(x interface{}) {
	*h = append(*h, x.([]int))
}

func (h *Heap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n - 1]
	*h = old[0 : n - 1]
	return x
}

func findScore(nums []int) int64 {
	result := int64(0)
	marked := make([]bool, len(nums))

	pq := &Heap{}
	heap.Init(pq)

	for i, num := range nums {
		heap.Push(pq, []int{num, i})
	}

	for pq.Len() > 0 {
		element := heap.Pop(pq).([]int)
		num := element[0]
		index := element[1]

		if !marked[index] {
			result += int64(num)
			marked[index] = true

			if index - 1 >= 0 {
				marked[index - 1] = true
			}

			if index + 1 < len(nums) {
				marked[index + 1] = true
			}
		}
	}

	return result
}

func main() {
	// result: 7
	// nums := []int{2,1,3,4,5,2}

	// result: 5
	nums := []int{2,3,5,1,3,2}

	// result: 
	// nums := []int{}

	result := findScore(nums)
	fmt.Printf("result = %v\n", result)
}

