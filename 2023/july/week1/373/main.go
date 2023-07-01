package main
import (
	"fmt"
	heap "container/heap"
)

type Pair struct {
	Sum int
	I int
	J int
}

type Heap []Pair

func (h Heap) Len() int			  { return len(h) }
func (h Heap) Less(a, b int) bool { return h[a].Sum < h[b].Sum }
func (h Heap) Swap(a, b int)	  { h[a], h[b] = h[b], h[a] }

func (h *Heap) Push(x interface{}) {
    *h = append(*h, x.(Pair))
}

func (h *Heap) Pop() interface{} {
    old := *h
    n := len(old)
    x := old[n - 1]
    *h = old[0 : n - 1]
    return x
}

func kSmallestPairs(nums1 []int, nums2 []int, k int) [][]int {
	m := len(nums1)
	n := len(nums2)
	result := make([][]int, 0, m * n)
	visited := make(map[string]bool)

    priorityQueue := &Heap{}
    heap.Init(priorityQueue)
	heap.Push(priorityQueue, Pair{Sum: nums1[0] + nums2[0], I: 0, J: 0})
	visited[fmt.Sprintf("%v,%v", 0, 0)] = true

	for k > 0 && priorityQueue.Len() > 0 {
		top := heap.Pop(priorityQueue).(Pair)
		i := top.I
		j := top.J
		result = append(result, []int{nums1[i], nums2[j]})

		key1 := fmt.Sprintf("%v,%v", i + 1, j)
		if i + 1 < m && !visited[key1] {
			heap.Push(priorityQueue, Pair{Sum: nums1[i + 1] + nums2[j], I: i + 1, J: j})
			visited[key1] = true
		}

		key2 := fmt.Sprintf("%v,%v", i, j + 1)
		if j + 1 < n && !visited[key2] {
			heap.Push(priorityQueue, Pair{Sum: nums1[i] + nums2[j + 1], I: i, J: j + 1})
			visited[key2] = true
		}

		k--
	}

	return result
}

func main() {
	// result: [[1,2],[1,4],[1,6]]
	// nums1 := []int{1,7,11}
	// nums2 := []int{2,4,6}
	// k := int(3)

	// result: [[1,1],[1,1]]
	// nums1 := []int{1,1,2}
	// nums2 := []int{1,2,3}
	// k := int(2)

	// result: [[1,3],[2,3]]
	nums1 := []int{1,2}
	nums2 := []int{3}
	k := int(3)

	// result: 
	// nums1 := []int{}
	// nums2 := []int{}
	// k := int(0)

	result := kSmallestPairs(nums1, nums2, k)
	fmt.Printf("result = %v\n", result)
}

