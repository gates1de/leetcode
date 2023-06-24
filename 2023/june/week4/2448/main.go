package main
import (
	"fmt"
)

func minCost(nums []int, cost []int) int64 {
	result := getCost(int64(nums[0]), nums, cost)
	left := int64(1000001)
	right := int64(0)
	for _, num := range nums {
		left = min(left, int64(num))
		right = max(right, int64(num))
	}

	for left < right {
		mid := int64((right + left) / 2)
		cost1 := getCost(mid, nums, cost)
		cost2 := getCost(mid + 1, nums, cost)
		result = min(cost1, cost2)
		if cost1 > cost2 {
			left = mid + 1
		} else {
			right = mid
		}
	}

	return result
}

func getCost(base int64, nums []int, cost []int) int64 {
	result := int64(0)
	for i := 0; i < len(nums); i++ {
		result += abs(int64(nums[i]), base) * int64(cost[i])
	}
	return result
}

func abs(a, b int64) int64 {
	if b > a {
		return b - a
	}
	return a - b
}

func max(a, b int64) int64 {
	if b > a {
		return b
	}
	return a
}

func min(a, b int64) int64 {
	if b < a {
		return b
	}
	return a
}

func main() {
	// result: 8
	// nums := []int{1,3,5,2}
	// cost := []int{2,3,1,14}

	// result: 0
	nums := []int{2,2,2,2,2}
	cost := []int{4,2,8,1,3}

	// result: 
	// nums := []int{}
	// cost := []int{}

	result := minCost(nums, cost)
	fmt.Printf("result = %v\n", result)
}

