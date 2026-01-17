package main

import (
	"fmt"
	"sort"
)

const modulo = int64(1e9 + 7)

func maximizeSquareArea(m int, n int, hFences []int, vFences []int) int {
	hEdges := getEdges(hFences, m)
	vEdges := getEdges(vFences, n)

	result := int64(0)
	for edge := range hEdges {
		if _, exists := vEdges[edge]; exists {
			if int64(edge) > result {
				result = int64(edge)
			}
		}
	}

	if result == 0 {
		return -1
	}

	return int((result * result) % modulo)
}

func getEdges(fences []int, border int) map[int]bool {
	set := make(map[int]bool)
	list := make([]int, len(fences))

	copy(list, fences)
	list = append(list, 1, border)
	sort.Ints(list)

	for i := 0; i < len(list); i++ {
		for j := i + 1; j < len(list); j++ {
			set[list[j] - list[i]] = true
		}
	}

	return set
}

func main() {
	// result: 4
	// m := int(4)
	// n := int(3)
	// hFences := []int{2,3}
	// vFences := []int{2}

	// result: -1
	m := int(6)
	n := int(7)
	hFences := []int{2}
	vFences := []int{4}

	// result: 
	// m := int(0)
	// n := int(0)
	// hFences := []int{}
	// vFences := []int{}

	result := maximizeSquareArea(m, n, hFences, vFences)
	fmt.Printf("result = %v\n", result)
}

