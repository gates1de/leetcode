package main
import (
	"fmt"
)

func maxSumOfThreeSubarrays(nums []int, k int) []int {
	bestSingleStart := int(0)
	bestDoubleStart := []int{0, k}
	result := []int{0, k, k * 2}
	currentWindowSumSingle := int(0)
	for i := 0; i < k; i++ {
		currentWindowSumSingle += nums[i]
	}

	currentWindowSumDouble := int(0)
	for i := k; i < k * 2; i++ {
		currentWindowSumDouble += nums[i]
	}

	currentWindowSumTriple := int(0)
	for i := k * 2; i < k * 3; i++ {
		currentWindowSumTriple += nums[i]
	}
	
	bestSingleSum := currentWindowSumSingle
	bestDoubleSum := currentWindowSumSingle + currentWindowSumDouble
	bestTripleSum := currentWindowSumSingle + currentWindowSumDouble + currentWindowSumTriple
	singleStartIndex := int(1)
	doubleStartIndex := k + 1
	tripleStartIndex := k * 2 + 1

	for tripleStartIndex <= len(nums) - k {
		currentWindowSumSingle = currentWindowSumSingle - nums[singleStartIndex - 1] + nums[singleStartIndex + k - 1]
		currentWindowSumDouble = currentWindowSumDouble - nums[doubleStartIndex - 1] + nums[doubleStartIndex + k - 1]
		currentWindowSumTriple = currentWindowSumTriple - nums[tripleStartIndex - 1] + nums[tripleStartIndex + k - 1]
	
		if currentWindowSumSingle > bestSingleSum {
			bestSingleStart = singleStartIndex
			bestSingleSum = currentWindowSumSingle
		}
	
		if currentWindowSumDouble + bestSingleSum > bestDoubleSum {
			bestDoubleStart[0] = bestSingleStart
			bestDoubleStart[1] = doubleStartIndex
			bestDoubleSum = currentWindowSumDouble + bestSingleSum
		}

		if currentWindowSumTriple + bestDoubleSum > bestTripleSum {
			result[0] = bestDoubleStart[0]
			result[1] = bestDoubleStart[1]
			result[2] = tripleStartIndex
			bestTripleSum = currentWindowSumTriple + bestDoubleSum
		}

		singleStartIndex += 1
		doubleStartIndex += 1
		tripleStartIndex += 1
	}
	
	return result
}

func main() {
	// result: [0,3,5]
	// nums := []int{1,2,1,2,6,7,5,1}
	// k := int(2)

	// result: [0,2,4]
	nums := []int{1,2,1,2,1,2,1,2,1}
	k := int(2)

	// result: []
	// nums := []int{}
	// k := int(0)

	result := maxSumOfThreeSubarrays(nums, k)
	fmt.Printf("result = %v\n", result)
}

