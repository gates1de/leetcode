package main
import (
	"fmt"
)

func increasingTriplet(nums []int) bool {
	n := len(nums)
	if n < 3 {
		return false
	}

	m := map[int]bool{}
	for _, num := range nums {
		m[num] = true
	}
	if len(m) < 3 {
		return false
	}

	for i := 1; i < n - 1; i++ {
		num := nums[i]
		leftIndex := i - 1
		rightIndex := i + 1
		for leftIndex >= 0 && rightIndex < n {
			left := nums[leftIndex]
			right := nums[rightIndex]

			if num <= left {
				leftIndex--
				continue
			}
			if right <= num {
				rightIndex++
				continue
			}

			return true
		}
	}

	return false
}

func main() {
	// result: true
	// nums := []int{1,2,3,4,5}

	// result: false
	// nums := []int{5,4,3,2,1}

	// result: true
	nums := []int{2,1,5,0,4,6}

	// result: 
	// nums := []int{}

	result := increasingTriplet(nums)
	fmt.Printf("result = %v\n", result)
}

