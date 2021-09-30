package main
import (
	"fmt"
	"sort"
)

func canPartitionKSubsets(nums []int, k int) bool {
    if k == 1 {
		return true
	}
    sum := 0
    max := 0
    for _, v := range nums {
        sum += v
        if v > max {
			max = v
		}
    }
    if sum % k != 0 || max > sum / k {
		return false
	}
    sort.Slice(nums, func(a, b int) bool { return nums[a] > nums[b] })
    seen := make([]bool, len(nums))
    return dfs(nums, seen, 0, 0, k, sum / k)
}

func dfs(nums []int, seen []bool, start, sum, k, target int) bool {
    if k == 1 {
		return true
	}
    if sum == target {
		return dfs(nums, seen, 0, 0, k - 1, target)
	}

    for i := start; i < len(nums); i++ {
        if !seen[i] && sum + nums[i] <= target {
            seen[i] = true
            if dfs(nums, seen, i + 1, nums[i] + sum, k, target) {
				return true
			}
            seen[i] = false
        }
    }
    return false
}

// Wrong Answer
func ngSolution(nums []int, k int) bool {
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

