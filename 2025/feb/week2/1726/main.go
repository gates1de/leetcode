package main
import (
	"fmt"
)

func tupleSameProduct(nums []int) int {
	n := len(nums)
	freq := make(map[int]int)
	result := int(0)

	for i, _ := range nums {
		for j := i + 1; j < n; j++ {
			freq[nums[i] * nums[j]]++
		}
	}

	for _, count := range freq {
		pairsOfEqualProduct := ((count - 1) * count) / 2
		result += 8 * pairsOfEqualProduct
	}

	return result
}

func main() {
	// result: 8
	// nums := []int{2,3,4,6}

	// result: 16
	nums := []int{1,2,4,5,10}

	// result: 
	// nums := []int{}

	result := tupleSameProduct(nums)
	fmt.Printf("result = %v\n", result)
}

