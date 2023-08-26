package main
import (
	"fmt"
	"math"
	"sort"
)

func findLongestChain(pairs [][]int) int {
	sort.Slice(pairs, func(a, b int) bool {
		return pairs[a][1] < pairs[b][1]
	})

	curr := math.MinInt32
	result := int(0)

	for _, pair := range pairs {
		if pair[0] > curr {
			result++
			curr = pair[1]
		}
	}

	return result
}

func main() {
	// result: 2
	// pairs := [][]int{{1,2},{2,3},{3,4}}

	// result: 3
	pairs := [][]int{{1,2},{7,8},{4,5}}

	// result: 
	// pairs := [][]int{}

	result := findLongestChain(pairs)
	fmt.Printf("result = %v\n", result)
}

