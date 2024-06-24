package main
import (
	"fmt"
)

func minKBitFlips(nums []int, k int) int {
	currentFlips := int(0)
	result := int(0)

	for i, num := range nums {
		if i >= k && nums[i - k] == 2 {
			currentFlips--
		}

		if currentFlips % 2 == num {
			if i + k > len(nums) {
				return -1
			}

			nums[i] = 2
			currentFlips++
			result++
		}
	}

	return result
}

func main() {
	// result: 2
	// nums := []int{0,1,0}
	// k := int(1)

	// result: -1
	// nums := []int{1,1,0}
	// k := int(2)

	// result: 3
	nums := []int{0,0,0,1,0,1,1,0}
	k := int(3)

	// result: 
	// nums := []int{}
	// k := int(0)

	result := minKBitFlips(nums, k)
	fmt.Printf("result = %v\n", result)
}

