package main
import (
	"fmt"
	heap "container/heap"
	"sort"
)

type Heap []int64

func (h Heap) Len() int           { return len(h) }
func (h Heap) Less(a, b int) bool { return h[a] < h[b] }
func (h Heap) Swap(a, b int)      { h[a], h[b] = h[b], h[a] }

func (h *Heap) Push(x interface{}) {
    *h = append(*h, x.(int64))
}

func (h *Heap) Pop() interface{} {
    old := *h
    n := len(old)
    x := old[n - 1]
    *h = old[0 : n - 1]
    return x
}

func maxScore(nums1 []int, nums2 []int, k int) int64 {
	n := len(nums1)
	pairs := make([][2]int, n)
	for i := 0; i < n; i++ {
		pairs[i] = [2]int{nums1[i], nums2[i]}
	}
	sort.Slice(pairs, func(a, b int) bool {
		return pairs[a][1] > pairs[b][1]
	})

    topKHeap := &Heap{}
    heap.Init(topKHeap)
	topKSum := int64(0)
	for i := 0; i < k; i++ {
		topKSum += int64(pairs[i][0])
        heap.Push(topKHeap, int64(pairs[i][0]))
	}

	result := topKSum * int64(pairs[k - 1][1])
        
	for i := k; i < n; i++ {
		topKSum += int64(pairs[i][0]) - heap.Pop(topKHeap).(int64)
        heap.Push(topKHeap, int64(pairs[i][0]))

		result = max(result, topKSum * int64(pairs[i][1]))
	}

	return result
}

func max(a, b int64) int64 {
	if b > a {
		return b
	}
	return a
}

func main() {
	// result: 12
	// nums1 := []int{1,3,3,2}
	// nums2 := []int{2,1,3,4}
	// k := int(3)

	// result: 30
	nums1 := []int{4,2,3,1,1}
	nums2 := []int{7,5,10,9,6}
	k := int(1)

	// result: 
	// nums1 := []int{}
	// nums2 := []int{}
	// k := int(0)

	result := maxScore(nums1, nums2, k)
	fmt.Printf("result = %v\n", result)
}

