package main
import (
	"fmt"
	"math"
	"sort"
)

func minPairSum(nums []int) int {
	n := len(nums)
	sort.Ints(nums)
	result := math.MinInt32

	for i, left := range nums[:n / 2] {
		right := nums[n - i - 1]
		sum := left + right
		if result < sum {
			result = sum
		}
	}

	return result
}

func main() {
	// result: 7
	// nums := []int{3,5,2,3}

	// result: 8
	nums := []int{3,5,4,2,4,6}

	// result: 
	// nums := []int{}

	result := minPairSum(nums)
	fmt.Printf("result = %v\n", result)
}

