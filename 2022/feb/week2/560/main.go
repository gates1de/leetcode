package main
import (
	"fmt"
)

func subarraySum(nums []int, k int) int {
	result := int(0)
	sumsFreq := map[int][]int{}
	sum := int(0)
	sums := make([]int, len(nums))
	for i, num := range nums {
		sum += num
		if sumsFreq[sum] == nil {
			sumsFreq[sum] = []int{}
		}
		sumsFreq[sum] = append(sumsFreq[sum], i)
		sums[i] = sum
	}

	for i, sum := range sums {
		if sum == k {
			result++
		}

		if indices, ok := sumsFreq[sum - k]; ok {
			for _, index := range indices {
				if i > index {
					result++
				}
			}
		}
	}
    return result
}

func main() {
	// result: 2
	// nums := []int{1, 1, 1}
	// k := int(2)

	// result: 2
	// nums := []int{1, 2, 3}
	// k := int(3)

	// result: 4
	// nums := []int{-1, 1, 2, 1, 3}
	// k := int(3)

	// result: 1
	// nums := []int{1}
	// k := int(1)

	// result: 1
	nums := []int{1}
	k := int(0)

	// result: 
	// nums := []int{}
	// k := int(0)

	// result: 
	// nums := []int{}
	// k := int(0)

	result := subarraySum(nums, k)
	fmt.Printf("result = %v\n", result)
}

