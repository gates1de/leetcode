package main
import (
	"fmt"
	heap "container/heap"
)

type Worker struct {
	SectionId int
    Cost int
}

type Heap []Worker

func (h Heap) Len() int { return len(h) }
func (h Heap) Less(a, b int) bool {
	if h[a].Cost == h[b].Cost {
		return h[a].SectionId < h[b].SectionId
	}
	return h[a].Cost < h[b].Cost
}
func (h Heap) Swap(a, b int) { h[a], h[b] = h[b], h[a] }

func (h *Heap) Push(x interface{}) {
    *h = append(*h, x.(Worker))
}

func (h *Heap) Pop() interface{} {
    old := *h
    n := len(old)
    x := old[n - 1]
    *h = old[0 : n - 1]
    return x
}

func totalCost(costs []int, k int, candidates int) int64 {
	n := len(costs)
    priorityQueue := &Heap{}
    heap.Init(priorityQueue)

	for i := 0; i < candidates; i++ {
		heap.Push(priorityQueue, Worker{SectionId: 0, Cost: costs[i]})
	}
	for i := max(candidates, n - candidates); i < n; i++ {
		heap.Push(priorityQueue, Worker{SectionId: 1, Cost: costs[i]})
	}

	result := int64(0)
	nextHead := candidates
	nextTail := n - 1 - candidates

	for i := 0; i < k; i++ {
		current := heap.Pop(priorityQueue).(Worker)
		cost := current.Cost
		sectionId := current.SectionId
		result += int64(cost)

		if nextHead <= nextTail {
			if sectionId == 0 {
				heap.Push(priorityQueue, Worker{SectionId: 0, Cost: costs[nextHead]})
				nextHead++
			} else {
				heap.Push(priorityQueue, Worker{SectionId: 1, Cost: costs[nextTail]})
				nextTail--
			}
		}
	}

	return result
}

func max(a, b int) int {
	if b > a {
		return b
	}
	return a
}

func main() {
	// result: 11
	// costs := []int{17,12,10,2,7,2,11,20,8}
	// k := int(3)
	// candidates := int(4)

	// result: 4
	costs := []int{1,2,4,1}
	k := int(3)
	candidates := int(3)

	// result: 
	// costs := []int{}
	// k := int(0)
	// candidates := int(0)

	result := totalCost(costs, k, candidates)
	fmt.Printf("result = %v\n", result)
}

