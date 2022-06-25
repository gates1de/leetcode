package main
import (
	"fmt"
)

func checkPossibility(nums []int) bool {
	n := len(nums)
	if n == 0 {
		return false
	}

	copyNums := make([]int, n)
	copy(copyNums, nums)

	firstCheckResult := false
	isChanged := false
	TOP:
	for i := 1; i < n; i++ {
		num := nums[i]

		for j := i - 1; j >= 0; j-- {
			pre := nums[j]
			if pre > num {
				if isChanged {
					firstCheckResult = false
					break TOP
				}
				nums[j] = num
				isChanged = true
				firstCheckResult = true
			}
		}
	}

	if firstCheckResult {
		return true
	}

	isChanged = false
	for i := 1; i < n; i++ {
		num := copyNums[i]

		for j := i - 1; j >= 0; j-- {
			pre := copyNums[j]
			if pre > num {
				if isChanged {
					return false
				}
				copyNums[i] = pre
				num = pre
				isChanged = true
			}
		}
	}

	return true
}

func main() {
	// result: true
	// nums := []int{4, 2, 3}

	// result: false
	// nums := []int{4, 2, 1}

	// result: true
   	// nums := []int{10, 5, 6, 9, 11}

	// result: false
   	// nums := []int{3, 4, 2, 3}

	// result: true
 	// nums := []int{5, 7, 1, 8}

	// result: true
	// nums := []int{2, 3}

	// result: true
	nums := []int{3, 2}

	// result: 
	// nums := []int{}

	result := checkPossibility(nums)
	fmt.Printf("result = %v\n", result)
}

