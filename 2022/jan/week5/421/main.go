package main
import (
	"fmt"
)

func findMaximumXOR(nums []int) int {
	result := int(0)
	for i := 0; i < len(nums); i++ {
		n1 := nums[i]
		for j := i + 1; j < len(nums); j++ {
			n2 := nums[j]
			if result < n1 ^ n2 {
				result = n1 ^ n2
			}
		}
	}
	return result
}

func main() {
	// result: 28
	// nums := []int{3, 10, 5, 25, 2, 8}

	// result: 127
	// nums := []int{14,70,53,83,49,91,36,80,92,51,66,70}

	// result: 0
	// nums := []int{2}

	// result: 815
	nums := []int{3,4,9,8,700,7,1,6,4,812,7,9,4,6,1,3,9,2,8,1}

	// result: 
	// nums := []int{}

	result := findMaximumXOR(nums)
	fmt.Printf("result = %v\n", result)
}

