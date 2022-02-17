package main
import (
	"fmt"
	"sort"
)

func combinationSum(candidates []int, target int) [][]int {
	sort.Ints(candidates)
	result := [][]int{}
	helper(0, []int{}, target, candidates, &result)
	return result
}

func helper(sum int, nums []int, target int, candidates []int, result *[][]int) {
	if sum >= target {
		if sum == target {
			*result = append(*result, nums)
		}
		return
	}

	n := len(nums)
	for i, cand := range candidates {
		if sum + cand > target {
			break
		}

		newNums := make([]int, n + 1)
		copy(newNums, nums)
		newNums[n] = cand
		helper(sum + cand, newNums, target, candidates[i:], result)
	}
}

func main() {
	// result: [[2,2,3],[7]]
	// candidates := []int{2, 3, 6, 7}
	// target := int(7)

	// result: [[2,2,2,2],[2,3,3],[3,5]]
	// candidates := []int{2, 3, 5}
	// target := int(8)

	// result: []
	candidates := []int{2}
	target := int(1)

	// result: 
	// candidates := []int{}
	// target := int(0)

	result := combinationSum(candidates, target)
	fmt.Printf("result = %v\n", result)
}

