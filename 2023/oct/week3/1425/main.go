package main
import (
	"fmt"
	"container/list"
	"math"
)

func constrainedSubsetSum(nums []int, k int) int {
	n := len(nums)
	deque := list.New()
	dp := make([]int, n)
	for i := 0; i < n; i++ {
		if deque.Len() > 0 && i - deque.Front().Value.(int) > k {
			deque.Remove(deque.Front())
		}

		dp[i] = nums[i]
		if deque.Len() > 0 {
			dp[i] += dp[deque.Front().Value.(int)]
		}
		for deque.Len() > 0 && dp[deque.Back().Value.(int)] < dp[i] {
			deque.Remove(deque.Back())
		}

		if dp[i] > 0 {
			deque.PushBack(i)
		}
	}

	result := math.MinInt32
	for _, num := range dp {
		result = max(result, num)
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
	// result: 37
	// nums := []int{10,2,-10,5,20}
	// k := int(2)

	// result: -1
	// nums := []int{-1,-2,-3}
	// k := int(1)

	// result: 23
	nums := []int{10,-2,-10,-5,20}
	k := int(2)

	// result: 
	// nums := []int{}
	// k := int(0)

	result := constrainedSubsetSum(nums, k)
	fmt.Printf("result = %v\n", result)
}

