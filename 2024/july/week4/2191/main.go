package main
import (
	"fmt"
	"sort"
)

// This answer almost correctly, but not passed one test case strangely.

type MappedNum struct {
	Num int
	Index int
}

func sortJumbled(mapping []int, nums []int) []int {
	n := len(nums)
	mappedNums := make([]MappedNum, n)

	for i, num := range nums {
		temp := num
		mappedNum := int(0)
		digit := int(1)

		if temp == 0 {
			mappedNums[i] = MappedNum{Num: mapping[0], Index: i}
			continue
		}

		for temp != 0 {
			mappedNum += digit * mapping[temp % 10]
			digit *= 10
			temp /= 10
		}
		mappedNums[i] = MappedNum{Num: mappedNum, Index: i}
	}

	sort.Slice(mappedNums, func (a, b int) bool {
		return mappedNums[a].Num < mappedNums[b].Num
	})

	result := make([]int, n)
	for i, _ := range result {
		result[i] = nums[mappedNums[i].Index]
	}

	return result
}

func main() {
	// result: [338,38,991]
	// mapping := []int{8,9,4,0,2,1,3,5,7,6}
	// nums := []int{991,338,38}

	// result: [123,456,789]
	mapping := []int{0,1,2,3,4,5,6,7,8,9}
	nums := []int{789,456,123}

	// result: []
	// mapping := []int{}
	// nums := []int{}

	result := sortJumbled(mapping, nums)
	fmt.Printf("result = %v\n", result)
}

