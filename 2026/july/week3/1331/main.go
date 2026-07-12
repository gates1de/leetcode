package main

import (
	"fmt"
	"sort"
)

type Pair struct {
	Value int
	Index int
}

func arrayRankTransform(arr []int) []int {
	n := len(arr)

	pairs := make([]Pair, n)
	for i := range n {
		pairs[i] = Pair{Value: arr[i], Index: i}
	}

	sort.Slice(pairs, func(i, j int) bool {
		return pairs[i].Value < pairs[j].Value
	})

	result := make([]int, n)
	rank := int(1)
	if n == 0 {
		return result
	}

	result[pairs[0].Index] = rank
	for i := 1; i < n; i++ {
		if pairs[i].Value > pairs[i - 1].Value {
			rank++
		}

		result[pairs[i].Index] = rank
	}

	return result
}

func main() {
	// result: [4,1,2,3]
	// arr := []int{40,10,20,30}

	// result:  [1,1,1]
	// arr := []int{100,100,100}

	// result: [5,3,4,2,8,6,7,1,3]
	arr := []int{37,12,28,9,100,56,80,5,12}

	// result: 
	// arr := []int{}

	result := arrayRankTransform(arr)
	fmt.Printf("result = %v\n", result)
}
