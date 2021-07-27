package main
import (
	"fmt"
	"sort"
)

func threeSumClosest(nums []int, target int) int {
	result := nums[0] + nums[1] + nums[2]
	sort.Ints(nums)
	for i := 0; i < len(nums); i++ {
		start := i + 1
		end := len(nums) - 1
		for start < end {
			cur := nums[i] + nums[start] + nums[end]
			if cur == target {
				return target
			} else if cur < target {
				start++
			} else {
				end--
			}
			diff1 := abs(cur, target)
			diff2 := abs(result, target)
			if diff1 < diff2 {
				result = cur
			}
		}
	}
	return result
}

// Slow & Use more memory
func mySolution(nums []int, target int) int {
	sort.Slice(nums, func (a, b int) bool { return nums[a] < nums[b] })
	result := int(10001)
	for i := 0; i < len(nums) - 2; i++ {
		num1 := nums[i]
		for j := i + 1; j < len(nums) - 1; j++ {
			num2 := nums[j]
			for k := j + 1; k < len(nums); k++ {
				num3 := nums[k]
				// fmt.Printf("num1 = %v, num2 = %v, num3 = %v\n", num1, num2, num3)
				sum := num1 + num2 + num3
				if abs(result, target) >= abs(sum, target) {
					result = sum
				}
				if sum == target {
					return target
				}
			}
		}
	}
	return result
}

func abs(a, b int) int {
	if b > a {
		return b - a
	}
	return a - b
}

func main() {
	// result: 2
	// nums := []int{-1, 2, 1, -4}
	// target := int(1)

	// result: 23
	nums := []int{-1,2,1,-4,-3,28,17,1,5,23,1,9,-15,2,3,10,-91,5,32,8,-9,1}
	target := int(23)

	// result: 
	// nums := []int{}
	// target := int(0)

	result := threeSumClosest(nums, target)
	fmt.Printf("result = %v\n", result)
}

