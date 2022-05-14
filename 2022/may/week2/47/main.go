package main
import (
	"fmt"
)

func permuteUnique(nums []int) [][]int {
	resultMap := map[[8]int]bool{}
	helper(0, [8]int{}, nums, resultMap)

	n := len(nums)
	result := make([][]int, len(resultMap))
	i := int(0)
	for key, _ := range resultMap {
		values := make([]int, n)
		copy(values, key[:n])
		result[i] = values
		i++
	}

	return result
}

func helper(index int, current [8]int, nums []int, resultMap map[[8]int]bool) {
	if len(nums) == 0 {
		resultMap[current] = true
		return
	}

	for i, num := range nums {
		newNums := make([]int, len(nums) - 1)
		copy(newNums[:i], nums[:i])
		copy(newNums[i:], nums[i + 1:])
		newCurrent := current
		newCurrent[index] = num
		helper(index + 1, newCurrent, newNums, resultMap)
	}
}

func main() {
	// result: [[1,1,2],[1,2,1],[2,1,1]]
	// nums := []int{1, 1, 2}

	// result: [[1,2,3],[1,3,2],[2,1,3],[2,3,1],[3,1,2],[3,2,1]]
	nums := []int{1, 2, 3}

	// result: 
	// nums := []int{}

	result := permuteUnique(nums)
	fmt.Printf("result = %v\n", result)
}

