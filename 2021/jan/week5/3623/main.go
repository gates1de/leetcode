package main
import (
	"fmt"
	"sort"
)

func nextPermutation(nums []int) {
	if len(nums) < 2 {
		return
	}

	targetNums := nums
	max := nums[len(nums) - 1]
	for i := len(nums) - 1; i >= 0; i-- {
		num := nums[i]
		if max < num {
			max = num
		}
		if max > num {
			targetNums = nums[i:]
			break
		}
	}
	// fmt.Printf("targetNums = %v\n", targetNums)

	swapIndex := int(-1)
TOP:
	for i := len(targetNums) - 1; i >= 0; i-- {
		num := targetNums[i]
		for j := i - 1; j >= 0; j-- {
			n := targetNums[j]
			fmt.Printf("num = %v, n = %v\n", num, n)
			if num > n {
				targetNums[i] = n
				targetNums[j] = num
				swapIndex = j
				break TOP
			}
		}
	}

	if swapIndex >= 0 {
		sortTargetNums := targetNums[swapIndex + 1:]
		sort.Slice(sortTargetNums, func (i, j int) bool { return sortTargetNums[i] < sortTargetNums[j] })
	} else if swapIndex < 0 {
		sort.Slice(targetNums, func (i, j int) bool { return targetNums[i] < targetNums[j] })
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

