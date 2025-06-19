package main
import (
	"fmt"
	"sort"
)

func partitionArray(nums []int, k int) int {
	sort.Ints(nums)
	result := int(1)
	pre := nums[0]

	for _, num := range nums {
		if num - pre > k {
			result++
			pre = num
		}
	}

	return result
}

func main() {
	// result: 2
	// nums := []int{3,6,1,2,5}
	// k := int(2)

	// result: 2
	// nums := []int{1,2,3}
	// k := int(1)

	// result: 3
	nums := []int{2,2,4,5}
	k := int(0)

	// result: 
	// nums := []int{}
	// k := int(0)

	result := partitionArray(nums, k)
	fmt.Printf("result = %v\n", result)
}

