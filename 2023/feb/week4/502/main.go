package main
import (
	"fmt"
	"container/heap"
	"sort"
)

type IntHeap []int

func (h IntHeap) Len() int { return len(h) }

func (h IntHeap) Less(i, j int) bool {
	return h[i] > h[j]
}

func (h IntHeap) Swap(i, j int) { 
    h[i], h[j] = h[j], h[i]
}

func (h *IntHeap) Push(x interface{}) {
    *h = append(*h, x.(int))
}
func (h *IntHeap) Pop() interface{} { 
    n := h.Len()
    old := *h
    x := old[n - 1]
    *h = old[:n - 1]
    return x
}

type Project struct {
	Capital int
	Profit int
}

func findMaximizedCapital(k int, w int, profits []int, capital []int) int {
	n := len(profits)
	projects := make([]Project, n)
	for i, profit := range profits {
		projects[i] = Project{Capital: capital[i], Profit: profit}
	}
	sort.Slice(projects, func(a, b int) bool {
		return projects[a].Capital < projects[b].Capital
	})

	h := &IntHeap{}
	heap.Init(h)
	pointer := int(0)
	for i := 0; i < k; i++ {
		for pointer < n && projects[pointer].Capital <= w {
            heap.Push(h, projects[pointer].Profit)
			pointer++
		}

		if h.Len() == 0 {
			break
		}
		w += heap.Pop(h).(int)
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

