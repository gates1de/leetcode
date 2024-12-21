package main
import (
	"fmt"
	heap "container/heap"
)

type Element struct {
	Index int
	Num int
}

type Heap []Element

func (h Heap) Len() int { return len(h) }

func (h Heap) Less(a, b int) bool {
	if h[a].Num == h[b].Num {
		return h[a].Index < h[b].Index
	}
	return h[a].Num < h[b].Num
}

func (h Heap) Swap(a, b int) { h[a], h[b] = h[b], h[a] }

func (h *Heap) Push(x interface{}) {
	*h = append(*h, x.(Element))
}

func (h *Heap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n - 1]
	*h = old[0 : n - 1]
	return x
}

func getFinalState(nums []int, k int, multiplier int) []int {
	pq := &Heap{}
	heap.Init(pq)

	for i, num := range nums {
		heap.Push(pq, Element{Index: i, Num: num})
	}

	for k > 0 {
		minElement := heap.Pop(pq).(Element)
		fmt.Println(minElement)
		heap.Push(pq, Element{Index: minElement.Index, Num: minElement.Num * multiplier})
		k--
	}

	result := make([]int, len(nums))
	for pq.Len() > 0 {
		element := heap.Pop(pq).(Element)
		result[element.Index] = element.Num
	}

	return result
}

func main() {
	// result: [8,4,6,5,6]
	// nums := []int{2,1,3,5,6}
	// k := int(5)
	// multiplier := int(2)

	// result: [16,8]
	nums := []int{1,2}
	k := int(3)
	multiplier := int(4)

	// result: []
	// nums := []int{}
	// k := int(0)
	// multiplier := int(0)

	result := getFinalState(nums, k, multiplier)
	fmt.Printf("result = %v\n", result)
}

