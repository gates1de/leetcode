package main
import (
	"fmt"
	"sort"
)

func maxOperations(nums []int, k int) int {
	sort.Slice(nums, func(i, j int) bool { return nums[i] < nums[j] })
	result := int(0)
	i := int(0)
	j := int(len(nums) - 1)
	for i < j {
		leftNum := nums[i]
		rightNum := nums[j]

		if leftNum >= k {
			i++
			continue
		}

		if rightNum >= k {
			j--
			continue
		}

		if leftNum + rightNum == k {
			result += 1
			i++
			j--
		} else if leftNum + rightNum < k {
			i++
		} else {
			j--
		}
	}
	return result
}

func main() {
	// nums := []int{1, 2, 3, 4}
	// k := 5

	// nums := []int{3, 1, 3, 4, 3}
	// k := 6

	nums := []int{3,1,5,1,1,1,1,1,2,2,3,2,2}
	k := 1

	result := maxOperations(nums, k)
	fmt.Printf("result = %v\n", result)
}

