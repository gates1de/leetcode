package main
import (
	"fmt"
	"container/heap"
	"math"
)

type Heap []int

func (h Heap) Len() int           { return len(h) }
func (h Heap) Less(a, b int) bool { return h[a] < h[b] }
func (h Heap) Swap(a, b int)      { h[a], h[b] = h[b], h[a] }

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

type MaxHeap struct {
    Heap
}

func (h MaxHeap) Less(i, j int) bool { return h.Heap[i] > h.Heap[j] }

func minimumDeviation(nums []int) int {
    if len(nums) < 1 {
        return 0
    }

    minNum := math.MaxInt64
    numsHeap := &MaxHeap{}
    heap.Init(numsHeap)
    for _, num := range nums {
        if num % 2 == 1 {
            heap.Push(numsHeap, num * 2)
            minNum = min(minNum, num * 2)
        } else {
            heap.Push(numsHeap, num)
            minNum = min(minNum, num)
        }
    }

    result := math.MaxInt64
    for numsHeap.Len() > 0 {
        num := heap.Pop(numsHeap).(int)
        result = min(result, num - minNum)

        if num % 2 == 0 {
            heap.Push(numsHeap, num / 2)
            minNum = min(minNum, num / 2)
            continue
        }
        break
    }

    return result
}

func min(a, b int) int {
    if a < b {
        return a
    }
    return b
}

func main() {
	// result: 1
	// nums := []int{1,2,3,4}

	// result: 3
	// nums := []int{4,1,5,20,3}

	// result: 3
	nums := []int{2,10,8}

	// result: 
	// nums := []int{}

	result := minimumDeviation(nums)
	fmt.Printf("result = %v\n", result)
}

