package main
import (
	"fmt"
	"container/heap"
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

// Out of memory
func ngSolution(heights []int, bricks int, ladders int) int {
	dp := make([][]int, bricks + 1)
	for i, _ := range dp {
		dp[i] = make([]int, ladders + 1)
	}
	helper(0, heights, bricks, ladders, dp)

	result := int(0)
	for _, row := range dp {
		for _, val := range row {
			if result < val {
				result = val
			}
		}
	}

	return result
}

func helper(currentIndex int, heights []int, bricks int, ladders int, dp [][]int) {
	if currentIndex < dp[bricks][ladders] {
		return
	}
	dp[bricks][ladders] = currentIndex

	if currentIndex + 1 >= len(heights) {
		return
	}

	currentHeight := heights[currentIndex]
	nextHeight := heights[currentIndex + 1]
	diff := nextHeight - currentHeight

	if diff <= 0 {
		helper(currentIndex + 1, heights, bricks, ladders, dp)
		return
	}

	if bricks >= diff {
		helper(currentIndex + 1, heights, bricks - diff, ladders, dp)
	}
	if ladders > 0 {
		helper(currentIndex + 1, heights, bricks, ladders - 1, dp)
	}
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

