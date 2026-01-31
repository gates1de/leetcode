package main

import (
	"fmt"
	"math"
)

func minimumCost(source string, target string, original []byte, changed []byte, cost []int) int64 {
	result := int64(0)
	minCost := make([][]int64, 26)
	for i, _ := range minCost {
		minCost[i] = make([]int64, 26)
		for j, _ := range minCost[i] {
			minCost[i][j] = math.MaxInt32
		}
	}

	for i := 0; i < len(original); i++ {
		startChar := original[i] - 'a'
		endChar := changed[i] - 'a'
		minCost[startChar][endChar] = min(minCost[startChar][endChar], int64(cost[i]))
	}

	for k := 0; k < 26; k++ {
		for i := 0; i < 26; i++ {
			for j := 0; j < 26; j++ {
				minCost[i][j] = min(minCost[i][j], minCost[i][k] + minCost[k][j])
			}
		}
	}

	for i := 0; i < len(source); i++ {
		if source[i] == target[i] {
			continue
		}

		sourceChar := source[i] - 'a'
		targetChar := target[i] - 'a'

		if minCost[sourceChar][targetChar] >= math.MaxInt32 {
			return -1
		}

		result += minCost[sourceChar][targetChar]
	}

	return result
}

func min(a, b int64) int64 {
	if b < a {
		return b
	}
	return a
}

func main() {
	// result: 28
	// source := "abcd"
	// target := "acbe"
	// original := []byte{'a','b','c','c','e','d'}
	// changed := []byte{'b','c','b','e','b','e'}
	// cost := []int{2,5,5,1,2,20}

	// result: 12
	// source := "aaaa"
	// target := "bbbb"
	// original := []byte{'a','c'}
	// changed := []byte{'c','b'}
	// cost := []int{1,2}

	// result: -1
	source := "abcd"
	target := "abce"
	original := []byte{'a'}
	changed := []byte{'e'}
	cost := []int{10000}

	// result: 
	// source := ""
	// target := ""
	// original := []byte{}
	// changed := []byte{}
	// cost := []int{}

	result := minimumCost(source, target, original, changed, cost)
	fmt.Printf("result = %v\n", result)
}

