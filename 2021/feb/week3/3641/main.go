package main
import (
	localHeap "./heap"
	"container/heap"
	"fmt"
)

func kWeakestRows(mat [][]int, k int) []int {
	h := &localHeap.Heap{}
	heap.Init(h)
	for i, nums := range mat {
		oneCount := int(0)
		for _, num := range nums {
			if num == 1 {
				oneCount++
			}
		}
		heap.Push(h, []int{oneCount, i})
		// fmt.Printf("heap = %v\n", h)
	}
	// fmt.Printf("heap = %v\n", h)
	result := make([]int, k)
	for i := 0; i < k; i++ {
		nums := heap.Pop(h).([]int)
		result[i] = nums[1]
	}
	return result
}

func main() {
	// mat := [][]int{
	// 	{1, 1, 0, 0, 0},
	// 	{1, 1, 1, 1, 0},
	// 	{1, 0, 0, 0, 0},
	// 	{1, 1, 0, 0, 0},
	// 	{1, 1, 1, 1, 1},
	// }
	// k := 3

	mat := [][]int{
		{1, 0, 0, 0},
		{1, 1, 1, 1},
		{1, 0, 0, 0},
		{1, 0, 0, 0},
	}
	k := 2

	result := kWeakestRows(mat, k)
	fmt.Printf("result = %v\n", result)
}

