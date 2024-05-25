package main
import (
	"fmt"
	"sort"
)

func beautifulSubsets(nums []int, k int) int {
	freqMap := make(map[int]int)
	sort.Ints(nums)
	return count(nums, k, freqMap, 0) - 1
}

func count(nums []int, diff int, freqMap map[int]int, i int) int {
	if i == len(nums) {
		return 1
	}

	result := count(nums, diff, freqMap, i + 1)
	num := nums[i]
	if _, ok := freqMap[num - diff]; !ok {
		freqMap[num]++

		result += count(nums, diff, freqMap, i + 1)
		freqMap[num]--

		if freqMap[num] == 0 {
			delete(freqMap, num)
		}
	}

	return result
}

func main() {
	// result: 4
	// nums := []int{2,4,6}
	// k := int(2)

	// result: 1
	nums := []int{1}
	k := int(1)

	// result: 
	// nums := []int{}
	// k := int(0)

	result := beautifulSubsets(nums, k)
	fmt.Printf("result = %v\n", result)
}

