package main
import (
	"fmt"
	"math"
	"sort"
)

func minDifference(nums []int) int {
	n := len(nums)
	if n <= 4 {
		return 0
	}

	sort.Ints(nums)
	result := math.MaxInt32

	right := n - 4
	for left := 0; left < 4; left++ {
		result = min(result, nums[right] - nums[left])
		right++
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
	// result: 0
	// nums := []int{5,3,2,4}

	// result: 1
	// nums := []int{1,5,0,10,14}

	// result: 0
	nums := []int{3,100,20}

	// result: 
	// nums := []int{}

	result := minDifference(nums)
	fmt.Printf("result = %v\n", result)
}

