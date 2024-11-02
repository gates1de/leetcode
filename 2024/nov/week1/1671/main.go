package main
import (
	"fmt"
	"math"
)

func minimumMountainRemovals(nums []int) int {
	n := len(nums)
	numsList := make([]int, n)
	for i, num := range nums {
		numsList[i] = num
	}

	listLengths := getLongestIncreasingSubsequenceLength(numsList)
	reverse(numsList)
	reverseListLengths := getLongestIncreasingSubsequenceLength(numsList)
	reverse(listLengths)

	result := math.MaxInt32
	for i, length := range listLengths {
		if length > 1 && reverseListLengths[i] > 1 {
			result = min(result, n - length - reverseListLengths[i] + 1)
		}
	}

	return result
}

func getLongestIncreasingSubsequenceLength(nums []int) []int {
	n := len(nums)
	listLengths := make([]int, n)
	for i, _ := range listLengths {
		listLengths[i] = 1
	}

	list := make([]int, 1)
	list[0] = nums[0]

	for i := 1; i < n; i++ {
		index := lowerBound(list, nums[i])

		if index == len(list) {
			list = append(list, nums[i])
		} else {
			list[index] = nums[i]
		}

		listLengths[i] = len(list)
	}

	return listLengths
}

func lowerBound(nums []int, target int) int {
	left := int(0)
	right := len(nums) - 1

	for left <= right {
		mid := left + (right - left) / 2
		if nums[mid] >= target {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}

	return left
}

func reverse(nums []int) {
	for i := 0; i < len(nums) / 2; i++ {
		nums[i], nums[len(nums) - i - 1] = nums[len(nums) - i - 1], nums[i]
	}
}

func min(a, b int) int {
	if b < a {
		return b
	}
	return a
}

func main() {
	// result: 0
	// nums := []int{1,3,1}

	// result: 3
	nums := []int{2,1,1,5,6,2,3,1}

	// result: 
	// nums := []int{}

	result := minimumMountainRemovals(nums)
	fmt.Printf("result = %v\n", result)
}

