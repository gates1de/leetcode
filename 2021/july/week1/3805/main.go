package main
import (
	"fmt"
	// "container/heap"
	"sort"
)

// type IntHeap []int
// 
// func (h *IntHeap) Less(i, j int) bool { return (*h)[i] < (*h)[j] }
// func (h *IntHeap) Swap(i, j int)      { (*h)[i], (*h)[j] = (*h)[j], (*h)[i] }
// func (h *IntHeap) Len() int           { return len(*h) }
// func (h *IntHeap) Pop() (v interface{}) {
//     *h, v = (*h)[:h.Len()-1], (*h)[h.Len()-1]
//     return v
// }
// func (h *IntHeap) Peek() int { return (*h)[0] }
// func (h *IntHeap) Push(v interface{}) { *h = append(*h, v.(int)) }
// 
// func kthSmallest(matrix [][]int, k int) int {
// 	h := &IntHeap{}
//     heap.Init(h)
// 	for _, row := range matrix {
// 		for _, num := range row {
// 			heap.Push(h, num)
// 		}
// 	}
// 	count := int(1)
// 	for count < k {
// 		_ = heap.Pop(h)
// 		count++
// 	}
// 	return heap.Pop(h).(int)
// }

func kthSmallest(matrix [][]int, k int) int {
	n := len(matrix)
	arr := make([]int, n * n)
	for i, row := range matrix {
		for j, num := range row {
			seq := i * n + j
			arr[seq] = num
		}
	}
	sort.Slice(arr, func (a, b int) bool { return arr[a] < arr[b] })
	return arr[k - 1]
}

func main() {
	// result: 13
	// matrix := [][]int{{1, 5, 9}, {10, 11, 13}, {12, 13, 15}}
	// k := int(8)

	// result: -5
	// matrix := [][]int{{-5}}
	// k := int(1)

	// result: 10
	matrix := [][]int{{8, 9, 10, 45}, {17, 23, 91, 519}, {105, 738, 2910, 2910}, {17, 23, 91, 519}}
	k := int(3)

	// result: 
	// matrix := [][]int{}
	// k := int(0)

	result := kthSmallest(matrix, k)
	fmt.Printf("result = %v\n", result)
}

