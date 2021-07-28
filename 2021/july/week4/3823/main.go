package main
import (
	"fmt"
	"sort"
)

// REF: https://leetcode.com/problems/partition-array-into-disjoint-intervals/discuss/1030544/Go
func partitionDisjoint(nums []int) int {
    copyNums := nums
    res := 1
    max := nums[0]
    subMax := nums[0]
    for i := 1; i < len(nums); i++ {
        if copyNums[i] < max {
            max = subMax
            res = i + 1
        } else {
            if copyNums[i] > subMax {
                subMax = copyNums[i]
            }
        }
    }
    return res
}

// Wrong Answer
func ngSolution(nums []int) int {
	minIndex := int(0)
	maxIndex := len(nums) - 1
	sortedNums := make([]int, maxIndex - minIndex + 1)
	copy(sortedNums, nums[minIndex:maxIndex + 1])
	sort.Slice(sortedNums, func (a, b int) bool { return sortedNums[a] < sortedNums[b] })
	// fmt.Printf("sortedNums = %v\n", sortedNums)
	for minIndex < maxIndex {
		minIndex = findMinIndex(minIndex, sortedNums[0], nums)
		maxIndex = findMaxIndex(maxIndex, sortedNums[len(sortedNums) - 1], nums)
		// fmt.Printf("minIndex = %v, maxIndex = %v\n", minIndex, maxIndex)
		if minIndex >= maxIndex {
			min := nums[minIndex]
			for _, num := range nums[:minIndex] {
				if min < num {
					minIndex++
					break
				}
			}
			break
		}

		sortedNums = make([]int, maxIndex - minIndex - 1)
		if len(sortedNums) == 0 {
			minIndex++
			break
		}
		copy(sortedNums, nums[minIndex + 1:maxIndex])
		sort.Slice(sortedNums, func (a, b int) bool { return sortedNums[a] < sortedNums[b] })
		// fmt.Printf("sortedNums = %v\n", sortedNums)
	}
	return minIndex
}

func findMinIndex(startIndex int, min int, nums []int) int {
	result := int(-1)
	for i := startIndex; i < len(nums); i++ {
		if min == nums[i] {
			result = i
		}
	}
	return result
}

func findMaxIndex(startIndex int, max int, nums []int) int {
	result := int(-1)
	for i := startIndex; i >= 0; i-- {
		if max == nums[i] {
			result = i
		}
	}
	return result
}

func main() {
	// result: 3 (left = [5,0,3], right = [8,6])
	// nums := []int{5, 0, 3, 8, 6}

	// result: 4 (left = [1,1,1,0], right = [6,12])
	// nums := []int{1, 1, 1, 0, 6, 12}

	// result: 8
	// nums := []int{5, 0, 3, 8, 6, 0, 9, 1, 10}

	// result: 1
	nums := []int{0, 2, 0, 2}

	// result: 
	// nums := []int{}

	result := partitionDisjoint(nums)
	fmt.Printf("result = %v\n", result)
}

