package main
import (
	"fmt"
)

func mostCompetitive(nums []int, k int) []int {
	result := make([]int, k)
	minIndex := int(0)
	for i := 0; i < k; i++ {
		target := nums[minIndex:len(nums) - k + i + 1]
		fmt.Printf("i = %v, minIndex = %v, result = %v, target = %v\n", i, minIndex, result, target)

		minNum, index := getMinFromSlice(target)
		if index < 0 {
			continue
		}
		result[i] = minNum
		minIndex = minIndex + index + 1
	}
	return result
}

func min(a int, b int) int {
	if a < b {
		return b
	}
	return a
}

// return value is (min num, index)
func getMinFromSlice(nums []int) (int, int) {
	result := int(-1)
	index := int(-1)
	for i, n := range nums {
		if result == -1 || result > n {
			result = n
			index = i
		}
	}
	return result, index
}

func main() {
	// nums := []int{3, 5, 2, 6}
	// k := 2

	nums := []int{2, 4, 3, 3, 5, 4, 9, 6}
	k := 4

	result := mostCompetitive(nums, k)
	fmt.Printf("result = %v\n", result)
}

