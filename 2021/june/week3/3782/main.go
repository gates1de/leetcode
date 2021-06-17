package main
import (
	"fmt"
	"container/heap"
)

func numSubarrayBoundedMax(nums []int, left int, right int) int {
	copyNums := nums
	cursor := int(0)
	result := int(0)
	start := int(0)
	last := int(-1)
	for i := 0; i < len(copyNums); i++ {
        if copyNums[i] < left {
            cursor += i - start + 1
            cursor -= i - last
        } else if copyNums[i] >= left && copyNums[i] <= right {
            last = i
            cursor += i - start + 1
        } else {
            result += cursor
            cursor = 0
            start = i + 1
            last = i
        }
    }
	result += cursor
    return result
}


type MaxHeap []int

func (h *MaxHeap) Less(i, j int) bool { return (*h)[i] > (*h)[j] }
func (h *MaxHeap) Swap(i, j int)      { (*h)[i], (*h)[j] = (*h)[j], (*h)[i] }
func (h *MaxHeap) Len() int           { return len(*h) }
func (h *MaxHeap) Pop() (v interface{}) {
        *h, v = (*h)[:h.Len()-1], (*h)[h.Len()-1]
        return v
}
func (h *MaxHeap) Peek() int { return (*h)[0] }
func (h *MaxHeap) Push(v interface{}) { *h = append(*h, v.(int)) }

// Time Limit Exceeded

func ngSolution(nums []int, left int, right int) int {
	result := [][]int{}
	helper(&result, nums, left, right)
	fmt.Printf("result = %v\n", result)
	return len(result)
}

func helper(result *[][]int, nums []int, left int, right int) {
	for i, num := range nums {
		sub := new(MaxHeap)
		// sub := []int{}
		// sub = append(sub, n1)
        heap.Push(sub, num)
		currentMax := sub.Peek()
		fmt.Printf("1: sub = %v\n", sub)
		if currentMax > right {
			continue
		}
		if left <= currentMax {
			*result = append(*result, *sub)
		}
		for j := i + 1; j < len(nums); j++ {
			// sub = append(sub, nums[j])
			heap.Push(sub, nums[j])
			fmt.Printf("2: sub = %v\n", sub)
			currentMax = sub.Peek()
			if currentMax > right {
				continue
			}
			if left <= currentMax {
				*result = append(*result, *sub)
			}
		}
	}
}

func main() {
	// result: 
	// nums := []int{2, 1, 4, 3}
	// left := int(2)
	// right := int(3)

	// result: 22
	nums := []int{73,55,36,5,55,14,9,7,72,52}
	left := int(32)
	right := int(69)

	// result: 
	// nums := []int{}
	// left := int(0)
	// right := int(0)

	result := numSubarrayBoundedMax(nums, left, right)
	fmt.Printf("result = %v\n", result)
}

