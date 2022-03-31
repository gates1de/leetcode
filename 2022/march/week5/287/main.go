package main
import (
	"fmt"
	"sort"
)

func findDuplicate(nums []int) int {
	sort.Ints(nums)
	pre := nums[0]
	for _, num := range nums[1:] {
		if pre == num {
			return num
		}
		pre = num
	}
	return -1
}

func main() {
	// result: 2
	// nums := []int{1, 3, 4, 2, 2}

	// result: 3
	// nums := []int{3, 1, 3, 4, 2}

	// result: 1
	nums := []int{1, 1}

	// result: 
	// nums := []int{}

	result := findDuplicate(nums)
	fmt.Printf("result = %v\n", result)
}

