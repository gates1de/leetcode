package main
import (
	"fmt"
	"container/list"
)

func maxSlidingWindow(nums []int, k int) []int {
	deque := list.New()
	result := make([]int, 0, len(nums))

	for i := 0; i < k; i++ {
		for deque.Front() != nil && nums[i] >= nums[deque.Back().Value.(int)] {
			deque.Remove(deque.Back())
		}
		deque.PushBack(i)
	}

	result = append(result, nums[deque.Front().Value.(int)])

	for i := k; i < len(nums); i++ {
		if deque.Front() != nil && deque.Front().Value.(int) == i - k {
			deque.Remove(deque.Front())
		}

		for deque.Back() != nil && nums[i] >= nums[deque.Back().Value.(int)] {
			deque.Remove(deque.Back())
		}

		deque.PushBack(i)
		result = append(result, nums[deque.Front().Value.(int)])
	}

	return result
}

func main() {
	// result: [3,3,5,5,6,7]
	// nums := []int{1,3,-1,-3,5,3,6,7}
	// k := int(3)

	// result: [1]
	nums := []int{1}
	k := int(1)

	// result: 
	// nums := []int{}
	// k := int(0)

	result := maxSlidingWindow(nums, k)
	fmt.Printf("result = %v\n", result)
}

