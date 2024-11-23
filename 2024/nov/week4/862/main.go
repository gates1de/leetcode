package main
import (
	"fmt"
	"math"
)

func shortestSubarray(nums []int, k int) int {
	n := len(nums)
	prefixSums := make([]int, n + 1)

	for i := 1; i <= n; i++ {
		prefixSums[i] = prefixSums[i - 1] + nums[i - 1]
	}

	indicesDeque := make([]int, 0, n)
	result := math.MaxInt32

	for i := 0; i <= n; i++ {
		for len(indicesDeque) > 0 && prefixSums[i] - prefixSums[indicesDeque[0]] >= k {
			first := indicesDeque[0]
			indicesDeque = indicesDeque[1:]
			result = min(result, i - first)
		}

		for len(indicesDeque) > 0 && prefixSums[i] <= prefixSums[indicesDeque[len(indicesDeque) - 1]] {
			indicesDeque = indicesDeque[:len(indicesDeque) - 1]
		}

		indicesDeque = append(indicesDeque, i)
	}

	if result == math.MaxInt32 {
		return -1
	}

	return result
}

func min(a, b int) int {
	if b < a {
		return b
	}
	return a
}

func main() {
	// result: 1
	// nums := []int{1}
	// k := int(1)

	// result: -1
	// nums := []int{1,2}
	// k := int(4)

	// result: 3
	nums := []int{2,-1,2}
	k := int(3)

	// result: 
	// nums := []int{}
	// k := int(0)

	result := shortestSubarray(nums, k)
	fmt.Printf("result = %v\n", result)
}

