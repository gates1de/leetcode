package main
import (
	"fmt"
	heap "container/heap"
)

type MaxHeap []int

func (h MaxHeap) Len() int           { return len(h) }
func (h MaxHeap) Less(i, j int) bool { return h[i] > h[j] }
func (h MaxHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h MaxHeap) Peek() int          { return h[0] }

func (h *MaxHeap) Push(x interface{}) {
    *h = append(*h, x.(int))
}

func (h *MaxHeap) Pop() interface{} {
    old := *h
    n := len(old)
    x := old[n-1]
    *h = old[0 : n-1]
    return x
}

func furthestBuilding(heights []int, bricks int, ladders int) int {
    h := &MaxHeap{}
    heap.Init(h)

    size := len(heights)
    if size == 1 {
        return 0
    }

    for i := 1; i < size; i++ {
        diff := heights[i] - heights[i - 1]

        if diff > 0 {
            if bricks >= diff {
                bricks -= diff
                heap.Push(h, diff)
            } else {
                if ladders == 0 {
                    return i - 1
                }

                ladders--
                heap.Push(h, diff)
                pop := heap.Pop(h).(int)

                // use a ladder instead of most used bricks
                if pop != diff {
                    bricks += pop - diff
                }
            }
        }
    }

    return size - 1
}

func main() {
	// result: 4
	// heights := []int{4,2,7,6,9,14,12}
	// bricks := int(5)
	// ladders := int(1)

	// result: 7
	// heights := []int{4,12,2,7,3,18,20,3,19}
	// bricks := int(10)
	// ladders := int(2)

	// result: 3
	heights := []int{14,3,19,3}
	bricks := int(17)
	ladders := int(0)

	// result: 
	// heights := []int{}
	// bricks := int(0)
	// ladders := int(0)

	result := furthestBuilding(heights, bricks, ladders)
	fmt.Printf("result = %v\n", result)
}

