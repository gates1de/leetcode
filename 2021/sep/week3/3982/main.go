package main
import (
	"fmt"
)

func findMaxConsecutiveOnes(nums []int) int {
	result := int(0)
	count := int(0)
	for _, n := range nums {
		if n == 1 {
			count++
		} else {
			if result < count {
				result = count
			}
			count = 0
		}
	}

	if result < count {
		return count
	}
	return result
}

func main() {
	// result: 3
	// nums := []int{1, 1, 0, 1, 1, 1}

	// result: 2
	// nums := []int{1, 0, 1, 1, 0, 1}

	// result: 1
	// nums := []int{1}

	// result: 5
	nums := []int{1,0,1,0,1,1,0,1,0,1,0,1,0,0,1,1,1,1,1,0,1,1,0,1,0,1,0,1,0,1,0,1,0,1,0,1,0}

	// result: 
	// nums := []int{}

	result := findMaxConsecutiveOnes(nums)
	fmt.Printf("result = %v\n", result)
}

