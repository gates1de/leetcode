package main
import (
	"fmt"
)

func getAverages(nums []int, k int) []int {
	n := len(nums)
	if k == 0 || n == 0 {
		return nums
	}

	sum := int(0)
	for i := 0; i < n; i++ {
		if i > k * 2 {
			break
		}
		sum += nums[i]
	}

	result := make([]int, n)
	for i, _ := range nums {
		if i < k || i + k >= n {
			result[i] = -1
			continue
		}

		if i > k {
			sum -= nums[i - k - 1]
			sum += nums[i + k]
		}
		result[i] = sum / (k * 2 + 1)
	}

	return result
}

func main() {
	// result: [-1,-1,-1,5,4,4,-1,-1,-1]
	// nums := []int{7,4,3,9,1,8,5,2,6}
	// k := int(3)

	// result: [100000]
	// nums := []int{100000}
	// k := int(0)

	// result: [-1]
	nums := []int{8}
	k := int(100000)

	// result: 
	// nums := []int{}
	// k := int(0)

	result := getAverages(nums, k)
	fmt.Printf("result = %v\n", result)
}

