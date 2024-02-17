package main
import (
	"fmt"
	"sort"
)

func findLeastNumOfUniqueInts(arr []int, k int) int {
	if len(arr) <= k {
		return 0
	}

	counts := make(map[int]int, len(arr))
	for _, num := range arr {
		counts[num]++
	}

	sortedCountValues := make([]int, 0, len(counts))
	for _, count := range counts {
		sortedCountValues = append(sortedCountValues, count)
	}

	sort.Ints(sortedCountValues)
	for k > 0 {
		count := sortedCountValues[0]

		if count > k {
			return len(sortedCountValues)
		} else {
			sortedCountValues = sortedCountValues[1:]
			k -= count

			if len(sortedCountValues) == 0 {
				return 0
			}
		}
	}

	return len(sortedCountValues)
}

func main() {
	// result: 1
	// arr := []int{5,5,4}
	// k := int(1)

	// result: 2
	arr := []int{4,3,1,1,3,3,2}
	k := int(3)

	// result: 
	// arr := []int{}
	// k := int(0)

	result := findLeastNumOfUniqueInts(arr, k)
	fmt.Printf("result = %v\n", result)
}

