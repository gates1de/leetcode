package main
import (
	"fmt"
	"sort"
)

func findUnsortedSubarray(nums []int) int {
	if len(nums) <= 1 {
		return 0
	}
	sortedNums := make([]int, len(nums))
	copy(sortedNums, nums)
	sort.Slice(sortedNums, func (a, b int) bool { return sortedNums[a] < sortedNums[b] })
	start := int(-1)
	end := int(-1)
	j := len(nums) - 1
	for i, num := range nums {
		if start >= 0 && end >= 0 {
			break
		}

		sortedNum := sortedNums[i]
		reverseNum := nums[j]
		reverseSortedNum := sortedNums[j]

		if num != sortedNum && start == -1 {
			start = i
		}
		if reverseNum != reverseSortedNum && end == -1 {
			end = j
		}
		i++
		j--
	}
	if start == -1 && end == -1 {
		return 0
	}
	// fmt.Printf("start = %v, end = %v\n", start, end)
	return end - start + 1
}

func main() {
	// nums := []int{2, 6, 4, 8, 10, 9, 15}
	// nums := []int{1, 2, 3, 4}
	nums := []int{1}
	result := findUnsortedSubarray(nums)
	fmt.Printf("result = %v\n", result)
}

