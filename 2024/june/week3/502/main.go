package main
import (
	"fmt"
	heap "container/heap"
)

type Project struct {
    Profit int
    Capital int
}

type MaxHeap []Project

func (h MaxHeap) Len() int {
	return len(h)
}

func (h MaxHeap) Less(i, j int) bool {
	return h[i].Profit > h[j].Profit
}

func (h MaxHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *MaxHeap) Push(x interface{}) {
	*h = append(*h, x.(Project))
}

func (h *MaxHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

type MinHeap []Project

func (h MinHeap) Len() int {
	return len(h)
}

func (h MinHeap) Less(i, j int) bool {
    if h[i].Capital == h[j].Capital{
        return h[i].Profit > h[j].Profit
    }
	return h[i].Capital < h[j].Capital
}

func (h MinHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *MinHeap) Push(x interface{}) {
	*h = append(*h, x.(Project))
}

func (h *MinHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func findMaximizedCapital(k int, w int, profits []int, capital []int) int {
    maxHeap := &MaxHeap{}
    minHeap := &MinHeap{}

    for i, profit := range profits {
		heap.Push(minHeap, Project{Profit: profit, Capital: capital[i]})
    }

    for k > 0 {
        for minHeap.Len() > 0 && (*minHeap)[0].Capital <= w {
            popedItem := heap.Pop(minHeap).(Project)
            heap.Push(maxHeap, popedItem)
        }

        if maxHeap.Len() <= 0 {
            break
        }

        popMaxHeap := heap.Pop(maxHeap).(Project)
        w += popMaxHeap.Profit
        k--
    }

    return w
}

func main() {
	// result: 4
	// k := int(2)
	// w := int(0)
	// profits := []int{1,2,3}
	// capital := []int{0,1,1}

	// result: 6
	k := int(3)
	w := int(0)
	profits := []int{1,2,3}
	capital := []int{0,1,2}

	// result: 
	// k := int(0)
	// w := int(0)
	// profits := []int{}
	// capital := []int{}

	result := findMaximizedCapital(k, w, profits, capital)
	fmt.Printf("result = %v\n", result)
}

