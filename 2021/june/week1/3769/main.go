package main
import (
	"fmt"
	"sort"
)

func longestConsecutive(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	sort.Slice(nums, func (a, b int) bool { return nums[a] < nums[b] })
	result := int(0)
	count := int(1)
	prev := nums[0]
	for _, num := range nums[1:] {
		if prev == num {
			continue
		}
		if prev + 1 != num {
			if result < count {
				result = count
			}
			count = 0
		}
		prev = num
		count++
	}
	if result < count {
		result = count
	}
    return result
}

func main() {
	// result: 4
	// nums := []int{100, 4, 200, 1, 3, 2}

	// result: 9
	// nums := []int{0,3,7,2,5,8,4,6,0,1}

	// result: 3
	// nums := []int{-2, -1, 0, 2, 4, 4}

	// result: 9
	nums := []int{-2,-1,0,2,4,4,5,6,7,19,18,17,16,15,14,-3,-5,13,-6,13,-7,12,11,-7,-4}

	// result: 
	// nums := []int{}

	result := longestConsecutive(nums)
	fmt.Printf("result = %v\n", result)
}

