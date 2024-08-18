package main
import (
	"fmt"
	"sort"
)

func smallestDistancePair(nums []int, k int) int {
	n := len(nums)
	sort.Ints(nums)
	low := int(0)
	high := nums[n - 1] - nums[0]

	for low < high {
		mid := (low + high) / 2
		count := helper(nums, mid)

		if count < k {
			low = mid + 1
		} else {
			high = mid
		}
	}

	return low
}

func helper(nums []int, maxDistance int) int {
	n := len(nums)
	count := int(0)
	left := int(0)

	for right := 0; right < n; right++ {
		for nums[right] - nums[left] > maxDistance {
			left++
		}

		count += right - left
	}

	return count
}

func main() {
	// result: 0
	// nums := []int{1,3,1}
	// k := int(1)

	// result: 0
	// nums := []int{1,1,1}
	// k := int(2)

	// result: 5
	nums := []int{1,6,1}
	k := int(3)

	// result: 
	// nums := []int{}
	// k := int(0)

	result := smallestDistancePair(nums, k)
	fmt.Printf("result = %v\n", result)
}

