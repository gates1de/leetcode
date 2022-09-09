package main
import (
	"fmt"
	"math"
	"sort"
)

func numberOfWeakCharacters(properties [][]int) int {
	n := len(properties)
	if n <= 1 {
		return 0
	}

	sort.Slice(properties, func(a, b int) bool {
		pa := properties[a]
		pb := properties[b]
		if pa[0] == pb[0] {
			return pa[1] > pb[1]
		}
		return pa[0] < pb[0]
	})

	result := int(0)
	maxDef := math.MinInt32
	for i := n - 1; i >= 0; i-- {
		if properties[i][1] < maxDef {
			result++
		}
		maxDef = max(maxDef, properties[i][1])
	}

	return result
}

func max(a, b int) int {
	if b > a {
		return b
	}
	return a
}

func main() {
	// result: 0
	// properties := [][]int{{5,5},{6,3},{3,6}}

	// result: 1
	// properties := [][]int{{2,2},{3,3}}

	// result: 1
	properties := [][]int{{1,5},{10,4},{4,3}}

	// result: 
	// properties := [][]int{}

	result := numberOfWeakCharacters(properties)
	fmt.Printf("result = %v\n", result)
}

