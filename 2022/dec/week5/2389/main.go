package main
import (
	"fmt"
	"sort"
)

func answerQueries(nums []int, queries []int) []int {
	sum := int(0)
	sums := make([]int, len(nums))
	sort.Ints(nums)
	for i, num := range nums {
		sum += num
		sums[i] = sum
	}

	result := make([]int, len(queries))
	for i, query := range queries {
		insertIndex := sort.SearchInts(sums, query)
		if insertIndex < len(sums) && sums[insertIndex] == query {
			insertIndex++
		}
		result[i] = insertIndex
	}
	return result
}

func main() {
	// result: [2,3,4]
	// nums := []int{4,5,2,1}
	// queries := []int{3,10,21}

	// result: [0]
	nums := []int{2,3,4,5}
	queries := []int{1}

	// result: 
	// nums := []int{}
	// queries := []int{}

	result := answerQueries(nums, queries)
	fmt.Printf("result = %v\n", result)
}

