package main
import (
	"fmt"
	"sort"
)

func lexicographicallySmallestArray(nums []int, limit int) []int {
	sortedNums := make([]int, len(nums))
	copy(sortedNums, nums)
	sort.Ints(sortedNums)

	currentGroup := int(0)
	numToGroup := make(map[int]int)
	numToGroup[sortedNums[0]] = currentGroup

	groupToList := make(map[int][]int)
	groupToList[currentGroup] = []int{sortedNums[0]}

	for i := 1; i < len(nums); i++ {
		if abs(sortedNums[i] - sortedNums[i - 1]) > limit {
			currentGroup++
		}

		numToGroup[sortedNums[i]] = currentGroup

		if _, exists := groupToList[currentGroup]; !exists {
			groupToList[currentGroup] = make([]int, 0)
		}
		groupToList[currentGroup] = append(groupToList[currentGroup], sortedNums[i])
	}

	for i, num := range nums {
		group := numToGroup[num]
		nums[i] = groupToList[group][0]
		groupToList[group] = groupToList[group][1:]
	}

	return nums
}

func abs(num int) int {
	if num < 0 {
		return -num
	}
	return num
}

func main() {
	// result: [1,3,5,8,9]
	// nums := []int{1,5,3,9,8}
	// limit := int(2)

	// result: [1,6,7,18,1,2]
	// nums := []int{1,7,6,18,2,1}
	// limit := int(3)

	// result: [1,7,28,19,10]
	nums := []int{1,7,28,19,10}
	limit := int(3)

	// result: []
	// nums := []int{}
	// limit := int(0)

	result := lexicographicallySmallestArray(nums, limit)
	fmt.Printf("result = %v\n", result)
}

