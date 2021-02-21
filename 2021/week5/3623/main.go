package main
import (
	"fmt"
	"sort"
)

func nextPermutation(nums []int) {
	if len(nums) < 2 {
		return
	}

	swapIndex := int(-1)
TOP:
	for i := len(nums) - 1; i > 0; i-- {
		num := nums[i]
		for j := i - 1; j >= 0; j-- {
			n := nums[j]
			fmt.Printf("num = %v, n = %v\n", num, n)
			if num > n {
				nums[i] = n
				nums[j] = num
				swapIndex = j
				break TOP
			}
		}
	}

	if swapIndex >= 0 {
		targetNums := nums[swapIndex + 1:]
		// fmt.Printf("swapIndex = %v, targetNums = %v\n", swapIndex, targetNums)
		sort.Slice(targetNums, func (i, j int) bool { return targetNums[i] < targetNums[j] })
	} else if swapIndex < 0 {
		sort.Slice(nums, func (i, j int) bool { return nums[i] < nums[j] })
	}
}

func copySlice(target []int) []int {
    result := make([]int, len(target))
    copy(result, target)
    return result
}

func main() {
	// nums := []int{1, 2, 3}
	// nums := []int{3, 2, 1}
	// nums := []int{1, 1, 5}
	// nums := []int{1}
	// nums := []int{1, 2}
	// nums := []int{1, 3, 2}
	nums := []int{4, 2, 0, 2, 3, 2, 0}


	nextPermutation(nums)
	fmt.Printf("result = %v\n", nums)
}

