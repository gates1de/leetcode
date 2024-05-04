package main
import (
	"fmt"
)

func minOperations(nums []int, k int) int {
	finalXor := int(0)
	for _, num := range nums {
		finalXor = finalXor ^ num
	}

	result := int(0)

	for k > 0 || finalXor > 0 {
		if k % 2 != finalXor % 2 {
			result++
		}

		k /= 2
		finalXor /= 2
	}

	return result
}

func main() {
	// result: 2
	// nums := []int{2,1,3,4}
	// k := int(1)

	// result: 0
	nums := []int{2,0,2,0}
	k := int(0)

	// result: 
	// nums := []int{}
	// k := int(0)

	result := minOperations(nums, k)
	fmt.Printf("result = %v\n", result)
}

