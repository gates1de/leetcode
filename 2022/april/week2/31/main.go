package main
import (
	"fmt"
	"sort"
)

func nextPermutation(nums []int)  {
	n := len(nums)
	if n <= 1 {
		return
	}

	cursor := len(nums) - 2
	for cursor >= 0 {
		if sort.IntsAreSorted(nums[cursor:]) {
			// if nums[cursor:] elements is same value, move cursor
			if sum(nums[cursor:]) == len(nums[cursor:]) * nums[cursor] {
				cursor--
				continue
			}
			nums[cursor], nums[cursor + 1] = nums[cursor + 1], nums[cursor]
			sort.Ints(nums[cursor + 1:])
			return
		}

		if nums[cursor] >= nums[cursor + 1] {
			cursor--
			continue
		}

		for i := n - 1; i > cursor; i-- {
			if nums[cursor] < nums[i] {
				nums[cursor], nums[i] = nums[i], nums[cursor]
				sort.Ints(nums[cursor + 1:])
				break
			}
		}
		return
	}

	sort.Ints(nums)
}

func sum(nums []int) int {
	result := int(0)
	for _, num := range nums {
		result += num
	}
	return result
}

func main() {
	// result: [1,3,2]
	// nums := []int{1, 2, 3}

	// result: [1,2,3]
	// nums := []int{3, 2, 1}

	// result: [1,5,1]
	// nums := []int{1, 1, 5}

	// result: [100]
	// nums := []int{100}

	// result: [2,1]
	// nums := []int{1, 2}

	// result: [2,1,3]
	// nums := []int{1, 3, 2}

	// result: [4,2,0,3,0,2,2]
	// nums := []int{4, 2, 0, 2, 3, 2, 0}

	// result: [1,1,5]
	// nums := []int{5, 1, 1}

	// result: [2,3,1,2,2,2,4,5,7]
	// nums := []int{2,2,7,5,4,3,2,2,1}

	// result: [2,3,3,1,3]
	nums := []int{2, 3, 1, 3, 3}

	// result: 
	// nums := []int{}

	nextPermutation(nums)
	fmt.Printf("result = %v\n", nums)
}

