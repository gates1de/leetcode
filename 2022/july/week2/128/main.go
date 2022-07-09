package main
import (
	"fmt"
	"sort"
)

func longestConsecutive(nums []int) int {
	n := len(nums)
	if n <= 1 {
		return n
	}

	sort.Ints(nums)

	result := int(1)
	count := int(1)
	pre := nums[0]
	for i := 1; i < n; i++ {
		num := nums[i]
		if num == pre {
			continue
		}

		if num - pre == 1 {
			count++
			if result < count {
				result = count
			}
		} else {
			count = 1
		}

		pre = num
	}

	return result
}

func main() {
	// result: 4
	// nums := []int{100,4,200,1,3,2}

	// result: 9
	// nums := []int{0,3,7,2,5,8,4,6,0,1}

	// result: 4
 	// nums := []int{1,11,12,2,3,44,5,13,14,0}

	// result: 1
	// nums := []int{1}

	// result: 2
	// nums := []int{0,-1}

	// result: 3
	nums := []int{1,2,0,1}

	// result: 
	// nums := []int{}

	result := longestConsecutive(nums)
	fmt.Printf("result = %v\n", result)
}

