package main
import (
    localHeap "heap/heap"
    "container/heap"
    "fmt"
	"math"
)

func minimumDeviation(nums []int) int {
	if len(nums) < 1 {
		return 0
	}
	minNum := math.MaxInt64
	numsHeap := &localHeap.MaxHeap{}
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
	// nums := []int{1, 2, 3, 4}

	// result: 3
	// nums := []int{4, 1, 5, 20, 3}

	// result: 3
	// nums := []int{2, 10, 8}

	// result: 315
	nums := []int{399, 908, 648, 357, 693, 502, 331, 649, 596, 698}

	result := minimumDeviation(nums)
	fmt.Printf("result = %v\n", result)
}

