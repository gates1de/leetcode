package main
import (
	"fmt"
	heap "container/heap"
)

type Counter struct {
    Num int
    Count int
}

type Heap []Counter

func (h Heap) Len() int           { return len(h) }
func (h Heap) Less(a, b int) bool { return h[a].Count > h[b].Count }
func (h Heap) Swap(a, b int)      { h[a], h[b] = h[b], h[a] }

func (h *Heap) Push(x interface{}) {
    *h = append(*h, x.(Counter))
}

func (h *Heap) Pop() interface{} {
    old := *h
    n := len(old)
    x := old[n - 1]
    *h = old[0 : n - 1]
    return x
}

func topKFrequent(nums []int, k int) []int {
    if len(nums) == 1 {
        return nums
    }

    counts := map[int]int{}
    for _, num := range nums {
        counts[num]++
    }

    numsHeap := &Heap{}
    heap.Init(numsHeap)
    for num, count := range counts {
        heap.Push(numsHeap, Counter{num, count})
    }

    result := make([]int, k)
    for i := 0; i < k; i++ {
        result[i] = (heap.Pop(numsHeap).(Counter)).Num
    }
    return result
}

func main() {
	// result: [1,2]
	// nums := []int{1,1,1,2,2,3}
	// k := int(2)

	// result: [1]
	nums := []int{1}
	k := int(1)

	// result: 
	// nums := []int{}
	// k := int(0)

	result := topKFrequent(nums, k)
	fmt.Printf("result = %v\n", result)
}

