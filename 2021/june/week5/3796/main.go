package main
import (
	"fmt"
)

// REF: https://leetcode.com/problems/max-consecutive-ones-iii/discuss/436641/Go-O(N)-Sliding-Window-Solution
func longestOnes(nums []int, k int) int {
	i := int(0)
	j := int(0)
    for i < len(nums) {
        k -= 1 - nums[i]
        i++
        if k < 0 {
            k += 1 - nums[j]
            j++
        }
    }
    return i - j
}

// Too slow & Use more memory
func mySolution(nums []int, k int) int {
	conCount := int(0)
	oneCount := int(0)
	shortcuts := []int{}
	maxConCount := int(0)
	for _, num := range nums {
		if num == 0 {
			if conCount > 0 {
				shortcuts = append(shortcuts, conCount)
				if maxConCount < conCount {
					maxConCount = conCount
				}
			}
			conCount = 0
			shortcuts = append(shortcuts, 0)
		} else if num == 1 {
			conCount++
			oneCount++
		}
	}

	if conCount > 0 {
		shortcuts = append(shortcuts, conCount)
	}
	if maxConCount < conCount {
		maxConCount = conCount
	}

	// fmt.Printf("shortcuts = %v, max = %v\n", shortcuts, maxConCount)
	if oneCount + k >= len(nums) {
		return len(nums)
	}

	firstZeroIndex := int(-1)
	for i, num := range shortcuts {
		if num == 0 {
			firstZeroIndex = i
			break
		}
	}
	if k == 0 || firstZeroIndex < 0 {
		return maxConCount
	}
	return helper(firstZeroIndex, maxConCount, shortcuts, k)
}

func helper(zeroIndex int, maxConCount int, nums []int, k int) int {
	// fmt.Printf("zeroIndex = %v\n", zeroIndex)
	copyNums := make([]int, len(nums))
	copy(copyNums, nums)
	remainK := k
	nextZeroIndex := zeroIndex
	i := zeroIndex
	for remainK > 0 && i < len(copyNums) - 1 {
		// fmt.Printf("i = %v, remainK = %v\n", i, remainK)
		num := copyNums[i]
		if num == 0 {
			copyNums[i] = 1
			remainK--
			if nextZeroIndex == zeroIndex && i > zeroIndex{
				nextZeroIndex = i
			}
		}
		i++
	}

	// fmt.Printf("nums = %v\n", copyNums)
	result := int(0)
	conCount := int(0)
	for _, num := range copyNums {
		if num == 0 {
			if conCount > 0 {
				if result < conCount {
					result = conCount
				}
			}
			conCount = 0
		} else if num >= 1 {
			conCount += num
		}
	}
	if result < conCount {
		result = conCount
	}
	// fmt.Printf("current result = %v\n", result)

	if i >= len(copyNums) - 1 {
		return result
	}
	next := helper(nextZeroIndex, maxConCount, nums, k)
	if result < next {
		return next
	}
	return result
}

func main() {
	// result: 6
	// nums := []int{1,1,1,0,0,0,1,1,1,1,0}
	// k := int(2)

	// result: 10
	// nums := []int{0,0,1,1,0,0,1,1,1,0,1,1,0,0,0,1,1,1,1}
	// k := int(3)

	// result: 8
	// nums := []int{0,1,0,1,0,1,0,1,0,1,0,1,0,1,0,1,0,1,0,1,0,1,0,1,1,0,1,0,1,0,0,1,0,1,1,0,0,1,0,1,0,1,0,1}
	// k := int(3)

	// result: 3
	// nums := []int{0,0,1,1,1,0,0}
	// k := int(0)

	// result: 25
	nums := []int{1,0,0,0,1,1,0,0,1,1,0,0,0,0,0,0,1,1,1,1,0,1,0,1,1,1,1,1,1,0,1,0,1,0,0,1,1,0,1,1}
	k := int(8)

	// result: 
	// nums := []int{}
	// k := int(0)

	result := longestOnes(nums, k)
	fmt.Printf("result = %v\n", result)
}

