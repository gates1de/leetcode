package main
import (
	"fmt"
	"sort"
)

func combinationSum2(candidates []int, target int) [][]int {
	result := make([][]int, 0)
	sort.Ints(candidates)
	temp := make([]int, 0)
	backtrack(&result, &temp, candidates, target, 0)
	return result
}

func backtrack(result *[][]int, temp *[]int, candidates []int, totalLeft int, index int) {
	if totalLeft < 0 {
		return
	}

	if totalLeft == 0 {
		nums := make([]int, len(*temp))
		copy(nums, *temp)
		*result = append(*result, nums)
		return
	}

	for i := index; i < len(candidates) && totalLeft >= candidates[i]; i++ {
		if i > index && candidates[i] == candidates[i - 1] {
			continue
		}

		*temp = append(*temp, candidates[i])
		backtrack(result, temp, candidates, totalLeft - candidates[i], i + 1)
		*temp = (*temp)[:len(*temp) - 1]
	}
}

func main() {
	// result: [[1,1,6],[1,2,5],[1,7],[2,6]]
	// candidates := []int{10,1,2,7,6,1,5}
	// target := int(8)

	// result: [[1,2,2],[5]]
	candidates := []int{2,5,2,1,2}
	target := int(5)

	// result: []
	// candidates := []int{}
	// target := int(0)

	result := combinationSum2(candidates, target)
	fmt.Printf("result = %v\n", result)
}

