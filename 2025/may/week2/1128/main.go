package main
import (
	"fmt"
)

func numEquivDominoPairs(dominoes [][]int) int {
	result := int(0)
	counter := make([]int, 100)

	for _, d := range dominoes {
		if d[0] > d[1] {
			d[0], d[1] = d[1], d[0]
		}

		val := d[0] * 10 + d[1]
		result += counter[val]
		counter[val]++
	}

	return result
}

func main() {
	// result: 1
	// dominoes := [][]int{{1,2},{2,1},{3,4},{5,6}}

	// result: 3
	dominoes := [][]int{{1,2},{1,2},{1,1},{1,2},{2,2}}

	// result: 
	// dominoes := [][]int{}

	result := numEquivDominoPairs(dominoes)
	fmt.Printf("result = %v\n", result)
}

