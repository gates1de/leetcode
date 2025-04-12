package main
import (
	"fmt"
)

func minOperations(nums []int, k int) int {
	m := make(map[int]bool)

	for _, num := range nums {
		if num < k {
			return -1
		} else if num > k {
			m[num] = true
		}
	}

	return len(m)
}

func main() {
	// result: 2
	// nums := []int{5,2,5,4,5}
	// k := int(2)

	// result: -1
	// nums := []int{2,1,2}
	// k := int(2)

	// result: 4
	nums := []int{9,7,5,3}
	k := int(1)

	// result: 
	// nums := []int{}
	// k := int(0)

	result := minOperations(nums, k)
	fmt.Printf("result = %v\n", result)
}

