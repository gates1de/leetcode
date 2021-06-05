package main
import (
	"fmt"
	"container/heap"
	"sort"
)

type Eng struct {
	Speed int
	Eff int
}
type EngHeap []Eng

func (h *EngHeap) Less(i, j int) bool { return (*h)[i].Speed < (*h)[j].Speed }
func (h *EngHeap) Swap(i, j int)      { (*h)[i], (*h)[j] = (*h)[j], (*h)[i] }
func (h *EngHeap) Len() int           { return len(*h) }
func (h *EngHeap) Pop() (v interface{}) {
    *h, v = (*h)[:h.Len()-1], (*h)[h.Len()-1]
    return v
}
func (h *EngHeap) Push(v interface{}) { *h = append(*h, v.(Eng)) }

func maxPerformance(n int, speed []int, efficiency []int, k int) int {
	engHeap := &EngHeap{}
	heap.Init(engHeap)
	engineers := make([]Eng, n)
	for i, s := range speed {
		engineers[i] = Eng{Speed: s, Eff: efficiency[i]}
	}
	sort.Slice(engineers, func (a, b int) bool { return engineers[a].Eff > engineers[b].Eff })
	// fmt.Printf("engineers = %v\n", engineers)

	result := int(0)
	totalSpeed := int(0)
	for _, eng := range engineers {
		heap.Push(engHeap, eng)
		totalSpeed += eng.Speed
		if engHeap.Len() > k {
			totalSpeed -= heap.Pop(engHeap).(Eng).Speed
		}
		if result < totalSpeed * eng.Eff {
			result = totalSpeed * eng.Eff
		}
	}
	return result % (1e9 + 7)
}

func main() {
	// result: 60
	// n := int(6)
	// k := int(2)
	// speed := []int{2, 10, 3, 1, 5, 8}
	// efficiency := []int{5, 4, 3, 9, 7, 2}

	// result: 68
	// n := int(6)
	// k := int(3)
	// speed := []int{2, 10, 3, 1, 5, 8}
	// efficiency := []int{5, 4, 3, 9, 7, 2}

	// result: 72
	// n := int(6)
	// k := int(4)
	// speed := []int{2, 10, 3, 1, 5, 8}
	// efficiency := []int{5, 4, 3, 9, 7, 2}

	// result: 123
	n := int(8)
	k := int(4)
	speed := []int{2, 10, 3, 1, 5, 8, 100, 4}
	efficiency := []int{5, 4, 3, 9, 7, 2, 1, 6}

	// result: 
	// n := int(0)
	// k := int(0)
	// speed := []int{}
	// efficiency := []int{}

	// result: 
	// n := int(0)
	// k := int(0)
	// speed := []int{}
	// efficiency := []int{}

	result := maxPerformance(n, speed, efficiency, k)
	fmt.Printf("result = %v\n", result)
}

