package main
import (
	"fmt"
    localHeap "heap/heap"
    "container/heap"
)

func topKFrequent(nums []int, k int) []int {
	if len(nums) == 1 {
		return nums
	}

	counts := map[int]int{}
	for _, num := range nums {
		counts[num]++
	}

	numsHeap := &localHeap.Heap{}
	heap.Init(numsHeap)
	for num, count := range counts {
		heap.Push(numsHeap, localHeap.Counter{num, count})
	}

	result := make([]int, k)
	for i := 0; i < k; i++ {
		result[i] = (heap.Pop(numsHeap).(localHeap.Counter)).Num
	}
	return result
}

func main() {
	// result: [1, 2]
	// nums := []int{1, 1, 1, 2, 2, 3}
	// k := int(2)

	// result: [1]
	// nums := []int{1}
	// k := int(1)

	// result: [1, 2, 3, 4]
	nums := []int{1, 2, 3, 4}
	k := int(4)

	// result: 
	// nums := []int{}
	// k := int(0)

	result := topKFrequent(nums, k)
	fmt.Printf("result = %v\n", result)
}

