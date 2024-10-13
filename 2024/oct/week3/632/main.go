package main
import (
	"fmt"
	"math"
	"sort"
)

func smallestRange(nums [][]int) []int {
	merged := make([][]int, 0)

	for i, arr := range nums {
		for _, num := range arr {
			merged = append(merged, []int{num, i})
		}
	}

	sort.Slice(merged, func(a, b int) bool { return merged[a][0] < merged[b][0] })
	freq := make(map[int]int)
	left := int(0)
	count := int(0)
	start := int(0)
	end := math.MaxInt32

	for right := 0; right < len(merged); right++ {
		freq[merged[right][1]]++

		if freq[merged[right][1]] == 1 {
			count++
		}

		for count == len(nums) {
			currentRange := merged[right][0] - merged[left][0]
			if currentRange < end - start {
				start = merged[left][0]
				end = merged[right][0]
			}

			freq[merged[left][1]]--

			if freq[merged[left][1]] == 0 {
				count--
			}

			left++
		}
	}

	return []int{start, end}
}

func main() {
	// result: [20,24]
	// nums := [][]int{{4,10,15,24,26},{0,9,12,20},{5,18,22,30}}

	// result: [1,1]
	nums := [][]int{{1,2,3},{1,2,3},{1,2,3}}

	// result: []
	// nums := [][]int{}

	result := smallestRange(nums)
	fmt.Printf("result = %v\n", result)
}

