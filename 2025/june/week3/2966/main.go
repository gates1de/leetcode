package main
import (
	"fmt"
	"sort"
)

func divideArray(nums []int, k int) [][]int {
	n := len(nums)
	sort.Ints(nums)
	result := make([][]int, n / 3)
	for i, _ := range result {
		result[i] = make([]int, 3)
	}

	for i := 0; i < n; i += 3 {
		if nums[i + 2] - nums[i] > k {
			return [][]int{}
		}
		result[i / 3][0] = nums[i]
		result[i / 3][1] = nums[i + 1]
		result[i / 3][2] = nums[i + 2]
	}

	return result
}

func main() {
	// result: [[1,1,3],[3,4,5],[7,8,9]]
	// nums := []int{1,3,4,8,7,9,3,5,1}
	// k := int(2)

	// result: []
	nums := []int{1,3,3,2,7,3}
	k := int(3)

	// result: 
	// nums := []int{}
	// k := int(0)

	result := divideArray(nums, k)
	fmt.Printf("result = %v\n", result)
}

