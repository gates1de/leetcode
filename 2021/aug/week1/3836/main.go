package main
import (
	"fmt"
	"math"
	"sort"
)

func twoSum(nums []int, target int) []int {
    m := make(map[int]int, len(nums))

    for i, num := range nums {
        if idx, ok := m[target - num]; ok {
            return []int{idx, i}
        }
        m[num] = i
    }
    return []int{0, 0}
}

// Slow & Use more memory
func mySolution(nums []int, target int) []int {
	if len(nums) == 2 {
		return []int{0, 1}
	}

	sortedNums := make([]int, len(nums))
	copy(sortedNums, nums)
	sort.Ints(sortedNums)
	one := int(0)
	two := int(0)
	for i, num := range sortedNums {
		remain := target - num
		if contains(remain, sortedNums[i + 1:]) {
			one = num
			two = remain
			break
		}
	}


	result := []int{10001, 10001}
	for i, num := range nums {
		if one == num {
			result[0] = i
			nums[i] = math.MinInt32
			break
		}
	}
	for i, num := range nums {
		if two == num {
			result[1] = i
			break
		}
	}
	return result
}

// conditions: nums is sorted
func contains(target int, nums []int) bool {
	for _, num := range nums {
		if target < num {
			return false
		}
		if target == num {
			return true
		}
	}
	return false
}

func main() {
	// result: [0,1]
	// nums := []int{2, 7, 11, 15}
	// target := int(9)

	// result: [1,2]
	// nums := []int{3, 2, 4}
	// target := int(6)

	// result: [0,1]
	// nums := []int{3, 3}
	// target := int(6)

	// result: [0,2]
	nums := []int{3, 2, 3}
	target := int(6)

	// result: 
	// nums := []int{}
	// target := int(0)

	result := twoSum(nums, target)
	fmt.Printf("result = %v\n", result)
}

