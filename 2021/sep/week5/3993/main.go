package main
import (
	"fmt"
	"sort"
)

func canPartitionKSubsets(nums []int, k int) bool {
	s := sum(nums)
	if s % k != 0 {
		return false
	}

	avg := s / k
	sort.Sort(sort.Reverse(sort.IntSlice(nums)))

	for i := 0; i < len(nums); i++ {
		if nums[i] == 0 || nums[i] == avg {
			nums[i] = 0
			continue
		} else if nums[i] > avg {
			return false
		}

		for j := i + 1; j < len(nums); j++ {
			if nums[i] + nums[j] <= avg {
				nums[i] += nums[j]
				nums[j] = 0
				if nums[i] == avg {
					break
				}
			}
		}
		if nums[i] < avg {
			return false
		}
	}
    return true
}

func sum(nums []int) int {
	result := int(0)
	for _, num := range nums {
		result += num
	}
	return result
}

func main() {
	// result: true
	// nums := []int{4, 3, 2, 3, 5, 2, 1}
	// k := int(4)

	// result: false
	// nums := []int{1, 2, 3, 4}
	// k := int(3)

	// result: true
	// nums := []int{4,3,2,3,5,2,1,1,1,2}
	// k := int(4)

	// result: false
	nums := []int{1, 1, 2, 4}
	k := int(4)

	// result: 
	// nums := []int{}
	// k := int(0)

	result := canPartitionKSubsets(nums, k)
	fmt.Printf("result = %v\n", result)
}

