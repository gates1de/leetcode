package main
import (
	"fmt"
	"sort"
)

func findUnsortedSubarray(nums []int) int {
	sortedNums := make([]int, len(nums))
	copy(sortedNums, nums)
	sort.Ints(sortedNums)
	i := int(0)
	j := len(nums) - 1

	for i < j {
		left := nums[i]
		right := nums[j]

		if left != sortedNums[i] && right != sortedNums[j] {
			return j - i + 1
		}

		if left == sortedNums[i] {
			i++
		}
		if right == sortedNums[j] {
			j--
		}
	}

	return 0
}

func main() {
	// result: 5
	// nums := []int{2,6,4,8,10,9,15}

	// result: 0
	// nums := []int{1, 2, 3, 4}

	// result: 0
	// nums := []int{0}

	// result: 
	nums := []int{1, 20, 19, 18, 17, 16, 14, 15, 13, 12}

	// result: 
	// nums := []int{}

	result := findUnsortedSubarray(nums)
	fmt.Printf("result = %v\n", result)
}

