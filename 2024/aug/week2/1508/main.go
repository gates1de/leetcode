package main
import (
	"fmt"
	"math"
)

const modulo = int(1e9 + 7)

func rangeSum(nums []int, n int, left int, right int) int {
	result := (helper(nums, n, right) - helper(nums, n, left - 1)) % int64(modulo)
	return int((int(result) + modulo) % modulo)
}

func helper(nums []int, n int, k int) int64 {
	minSum := minInNums(nums)
	maxSum := sum(nums)
	left := minSum
	right := maxSum

	for left <= right {
		mid := left + (right - left) / 2
		count, _ := countAndSum(nums, n, mid)
		if count >= k {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	count, val := countAndSum(nums, n, left)
	return val - int64(left * (count - k))
}

func countAndSum(nums []int, n int, target int) (int, int64) {
	count := int(0)
	currentSum := int(0)
	totalSum := int64(0)
	windowSum := int(0)

	i := int(0)
	for j := 0; j < n; j++ {
		currentSum += nums[j]
		windowSum += nums[j] * (j - i + 1)
		for currentSum > target {
			windowSum -= currentSum
			currentSum -= nums[i]
			i++
		}

		count += j - i + 1
		totalSum += int64(windowSum)
	}

	return count, totalSum
}

func minInNums(nums []int) int {
	result := math.MaxInt32
	for _, num := range nums {
		if num < result {
			result = num
		}
	}
	return result
}

func sum(nums []int) int {
	result := int(0)
	for _, num := range nums {
		result += num
	}
	return result
}

func main() {
	// result: 13
	// nums := []int{1,2,3,4}
	// n := int(4)
	// left := int(1)
	// right := int(5)

	// result: 6
	// nums := []int{1,2,3,4}
	// n := int(4)
	// left := int(3)
	// right := int(4)

	// result: 50
	nums := []int{1,2,3,4}
	n := int(4)
	left := int(1)
	right := int(10)

	// result: 
	// nums := []int{}
	// n := int(0)
	// left := int(0)
	// right := int(0)

	result := rangeSum(nums, n, left, right)
	fmt.Printf("result = %v\n", result)
}

