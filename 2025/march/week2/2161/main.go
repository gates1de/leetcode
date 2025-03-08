package main
import (
	"fmt"
)

func pivotArray(nums []int, pivot int) []int {
	n := len(nums)
	result := make([]int, n)
	lessIndex := int(0)
	greaterIndex := n - 1

	i := int(0)
	j := n - 1
	for i < n {
		if nums[i] < pivot {
			result[lessIndex] = nums[i]
			lessIndex++
		}
		if nums[j] > pivot {
			result[greaterIndex] = nums[j]
			greaterIndex--
		}

		i++
		j--
	}

	for lessIndex <= greaterIndex {
		result[lessIndex] = pivot
		lessIndex++
	}

	return result
}

func main() {
	// result: [9,5,3,10,10,12,14]
	// nums := []int{9,12,5,10,14,3,10}
	// pivot := int(10)

	// result: [-3,2,4,3]
	nums := []int{-3,4,3,2}
	pivot := int(2)

	// result: []
	// nums := []int{}
	// pivot := int(0)

	result := pivotArray(nums, pivot)
	fmt.Printf("result = %v\n", result)
}

