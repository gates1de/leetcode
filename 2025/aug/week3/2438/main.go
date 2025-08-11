package main

import (
	"fmt"
)

const modulo = int(1e9 + 7)

func productQueries(n int, queries [][]int) []int {
	bins := make([]int, 0)
	rep := int(1)
	for n > 0 {
		if n % 2 == 1 {
			bins = append(bins, rep)
		}

		n /= 2
		rep *= 2
	}

	result := make([]int, 0, len(queries))
	for _, query := range queries {
		cur := int(1)
		for i := query[0]; i <= query[1]; i++ {
			cur = (cur * bins[i]) % modulo
		}
		result = append(result, cur)
	}

	return result
}

func main() {
	// result: [2,4,64]
	// n := int(15)
	// queries := [][]int{{0,1},{2,2},{0,3}}

	// result: [2]
	n := int(2)
	queries := [][]int{{0,0}}

	// result: []
	// n := int(0)
	// queries := [][]int{}

	result := productQueries(n, queries)
	fmt.Printf("result = %v\n", result)
}

