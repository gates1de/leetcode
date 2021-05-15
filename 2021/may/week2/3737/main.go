package main
import (
	"fmt"
	"container/heap"
)
type MaxHeap []int

func (h *MaxHeap) Less(i, j int) bool { return (*h)[i] > (*h)[j] }
func (h *MaxHeap) Swap(i, j int)      { (*h)[i], (*h)[j] = (*h)[j], (*h)[i] }
func (h *MaxHeap) Len() int           { return len(*h) }
func (h *MaxHeap) Pop() (v interface{}) {
	*h, v = (*h)[:h.Len()-1], (*h)[h.Len()-1]
	return v
}
func (h *MaxHeap) Push(v interface{}) { *h = append(*h, v.(int)) }

func isPossible(target []int) bool {
	maxHeap, sum := new(MaxHeap), 0
	for _, n := range target {
		heap.Push(maxHeap, n)
		sum += n
	}

	for {
		a := heap.Pop(maxHeap).(int)
		sum -= a
		if a == 1 || sum == 1 {
			return true
		}
		if sum >= a || sum == 0 || a%sum == 0 {
			return false
		}
		a %= sum
		heap.Push(maxHeap, a)
		sum += a
	}
}


// Slow & Use more memory
func myAnswer(target []int) bool {
	nums := copySlice(target)
	currentSum := sum(nums)
	if len(nums) == 1 {
		if target[0] == 1 {
			return true
		} else {
			return false
		}
	} else if len(nums) % 2 == 0 && currentSum % 2 == 0 {
		return false
	}
	return helper(nums, target)
}

func helper(nums []int, target []int) bool {
	currentSum := sum(nums)
	currentMaxIndex := int(0)
	currentMax := int(0)
	for i, num := range nums {
		if currentMax < num {
			currentMax = num
			currentMaxIndex = i
		}
	}
	prev := currentMax - (currentSum - currentMax)
	if prev < 1 {
		return false
	}
	otherSum := currentSum - currentMax
	if otherSum == 1 {
		return true
	} else if otherSum < prev {
		for prev > otherSum {
			prev -= otherSum
		}
	}
	nums[currentMaxIndex] = prev
	if sum(nums) == len(nums) {
		return true
	}
	// fmt.Printf("nums = %v, prev = %v, otherSum = %v\n", nums, prev, otherSum)
	return helper(nums, target)
}

func sum(nums []int) int {
	result := int(0)
	for _, num := range nums {
		result += num
	}
	return result
}

func getMax(nums []int) int {
	result := int(0)
	for _, num := range nums {
		if result < num {
			result = num
		}
	}
	return result
}

func copySlice(target []int) []int {
    result := make([]int, len(target))
    copy(result, target)
    return result
}

func main() {
	// result: true
	// target := []int{9, 3, 5}

	// result: false
	// target := []int{1, 1, 1, 2}

	// result: true
	// target := []int{8, 5}

	// result: false
	// target := []int{9, 9, 9}

	// result: true
	// target := []int{1, 1000000000}

	// result: true
	// target := []int{5, 2}

	// result: false
	// target := []int{2, 900000002}

	// result: false
	target := []int{5, 50}

	// result: 
	// target := []int{}

	result := isPossible(target)
	fmt.Printf("result = %v\n", result)
}

