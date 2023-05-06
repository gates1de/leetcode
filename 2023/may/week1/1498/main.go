package main
import (
	"fmt"
	"sort"
)

const modulo = int(1e9 + 7)

func numSubseq(nums []int, target int) int {
	n := len(nums)
	sort.Ints(nums)
	power := make([]int, n)
	power[0] = 1
	for i := 1; i < n; i++ {
		power[i] = (power[i - 1] * 2) % modulo
	}

	result := int(0)
	for left := 0; left < n; left++ {
		right := binarySearch(target - nums[left], nums) - 1
		if right >= left {
			result += power[right - left]
			result %= modulo
		}
	}

	return result
}

func binarySearch( target int, nums []int) int {
	left := int(0)
	right := len(nums) - 1

	for left <= right {
		mid := left + (right - left) / 2
		if nums[mid] <= target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	return left
}
    
func main() {
	// result: 4
	// nums := []int{3,5,6,7}
	// target := int(9)

	// result: 6
	// nums := []int{3,3,6,8}
	// target := int(10)

	// result: 61
	nums := []int{2,3,3,4,6,7}
	target := int(12)

	// result: 
	// nums := []int{}
	// target := int(0)

	result := numSubseq(nums, target)
	fmt.Printf("result = %v\n", result)
}

