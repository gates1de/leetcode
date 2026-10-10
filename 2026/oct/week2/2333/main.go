package main

import (
	"fmt"
	"sort"
)

func minSumSquareDiff(nums1 []int, nums2 []int, k1 int, k2 int) int64 {
	if len(nums1) == 0 {
		return 0
	}

	differences := make([]int, len(nums1))
	for i := range nums1 {
		difference := nums1[i] - nums2[i]
		if difference < 0 {
			difference = -difference
		}
		differences[i] = difference
	}
	sort.Ints(differences)

	remaining := int64(k1) + int64(k2)
	topCount := 1
	level := int64(differences[len(differences)-1])
	extra := int64(0)
	partial := false

	for i := len(differences) - 2; i >= 0; i-- {
		nextLevel := int64(differences[i])
		if nextLevel == level {
			topCount++
			continue
		}

		cost := (level - nextLevel) * int64(topCount)
		if remaining < cost {
			level -= remaining / int64(topCount)
			extra = remaining % int64(topCount)
			partial = true
			break
		}

		remaining -= cost
		level = nextLevel
		topCount++
		if level == 0 {
			break
		}
	}

	result := int64(0)
	for _, difference := range differences[:len(differences)-topCount] {
		value := int64(difference)
		result += value * value
	}

	if partial {
		remainingCount := int64(topCount) - extra
		result += remainingCount * level * level
		lowerLevel := level - 1
		result += extra * lowerLevel * lowerLevel
	} else {
		result += int64(topCount) * level * level
	}

	return result
}

func main() {
	// result: 579
	// nums1 := []int{1,2,3,4}
	// nums2 := []int{2,10,20,19}
	// k1 := int(0)
	// k2 := int(0)

	// result: 43
	nums1 := []int{1, 4, 10, 12}
	nums2 := []int{5, 8, 6, 9}
	k1 := int(1)
	k2 := int(1)

	// result: 0
	// nums1 := []int{}
	// nums2 := []int{}
	// k1 := int(0)
	// k2 := int(0)

	result := minSumSquareDiff(nums1, nums2, k1, k2)
	fmt.Printf("result = %v\n", result)
}
