package main
import (
	"fmt"
	"sort"
)

func subsetsWithDup(nums []int) [][]int {
	m := map[string]bool{}
	result := [][]int{[]int{}}
	sort.Ints(nums)
	for i, num := range nums {
		current := []int{num}
		helper(current, nums[i + 1:], m, &result)
	}
	return result
}

func helper(current []int, nums []int, m map[string]bool, result *[][]int) {
	sort.Ints(current)
	currentStr := ""
	comma := ""
	for _, num := range current {
		currentStr += comma + fmt.Sprintf("%v", num)
		comma = ","
	}
	if !m[currentStr] {
		*result = append(*result, current)
		m[currentStr] = true
	}

	if len(nums) == 0 {
		return
	}

	for i, num := range nums {
		newCurrent := append([]int{num}, current...)
		helper(newCurrent, nums[i + 1:], m, result)
	}
}

func remove(index int, nums []int) []int {
	if len(nums) == 0 {
		return []int{}
	}
	copyNums := make([]int, len(nums))
	copy(copyNums, nums)
	return append(copyNums[:index], copyNums[index + 1:]...)
}

func main() {
	// result: [[],[1],[1,2],[1,2,2],[2],[2,2]]
	// nums := []int{1, 2, 2}

	// result: [[],[0]]
	// nums := []int{0}

	// result: abbreviation
	nums := []int{1, 2, 3, 4, 10, 5, 6, 7, 8, 9}

	// result: 
	// nums := []int{}

	result := subsetsWithDup(nums)
	fmt.Printf("result = %v\n", result)
}

