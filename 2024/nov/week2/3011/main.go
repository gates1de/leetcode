package main
import (
	"fmt"
	"math/bits"
)

func canSortArray(nums []int) bool {
	n := len(nums)
	copyNums := make([]int, n)
	copy(copyNums, nums)

	for i := 0; i < n - 1; i++ {
		if copyNums[i] <= copyNums[i + 1] {
			continue
		}

		if bits.OnesCount32(uint32(copyNums[i])) == bits.OnesCount32(uint32(copyNums[i + 1])) {
			copyNums[i], copyNums[i + 1] = copyNums[i + 1], copyNums[i]
		} else {
			return false
		}
	}

	for i := n - 1; i >= 1; i-- {
		if copyNums[i] >= copyNums[i - 1] {
			continue
		}

		if bits.OnesCount32(uint32(copyNums[i])) == bits.OnesCount32(uint32(copyNums[i - 1])) {
			copyNums[i], copyNums[i - 1] = copyNums[i - 1], copyNums[i]
		} else {
			return false
		}
	}

	return true
}

func main() {
	// result: true
	// nums := []int{8,4,2,30,15}

	// result: true
	// nums := []int{1,2,3,4,5}

	// result: false
	nums := []int{3,16,8,4,2}

	// result: false
	// nums := []int{}

	result := canSortArray(nums)
	fmt.Printf("result = %v\n", result)
}

