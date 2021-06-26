package main
import (
	"fmt"
)

func countSmaller(nums []int) []int {
	index := make([]int, 0, len(nums))
	for i := range nums {
		index = append(index, i)
	}
	inversions := make([]int, len(nums), len(nums))
	merge(0, len(nums)-1, nums, index, inversions)
	return inversions
}

func merge(left, right int, nums, index, inversions []int) {
	if left >= right {
		return
	}
	mid := (left + right) / 2
	merge(left, mid, nums, index, inversions)
	merge(mid+1, right, nums, index, inversions)

	mergeHelper(left, right, nums, index, inversions)
}

func mergeHelper(left, right int, nums, index, inversions []int) {
	start, end := left, right
	mid := (left + right) / 2
	leftMid, rightMid := mid, mid+1

	newIndex := make([]int, 0)
	smaller := 0
	for left <= leftMid && rightMid <= right {
		if nums[index[left]] <= nums[index[rightMid]] {
			newIndex = append(newIndex, index[left])
			inversions[index[left]] += smaller
			left++
		} else {
			newIndex = append(newIndex, index[rightMid])
			rightMid++
			smaller++
		}
	}

	for left <= leftMid {
		inversions[index[left]] += smaller
		newIndex = append(newIndex, index[left])
		left++
	}
	for rightMid <= right {
		newIndex = append(newIndex, index[rightMid])
		rightMid++
	}

	copy(index[start:end+1], newIndex)
}

// Time Limit Exceeded
func ngSolution(nums []int) []int {
	result := make([]int, len(nums))
	for i, n1 := range nums {
		for _, n2 := range nums[i + 1:] {
			if n1 > n2 {
				result[i]++
			}
		}
	}
	return result
}

func main() {
	// result: [2,1,1,0]
	nums := []int{5, 2, 6, 1}

	// result: [0]
	// nums := []int{-1}

	// result: [0, 0]
	// nums := []int{-1, -1}

	// result: 
	// nums := []int{}

	result := countSmaller(nums)
	fmt.Printf("result = %v\n", result)
}

