package main

import (
	"fmt"
)

func closestTarget(words []string, target string, startIndex int) int {
	result := len(words)
	n := len(words)
	for i, word := range words {
		if word == target {
			dist := abs(i - startIndex)
			result = min(result, min(dist, n - dist))
		}
	}

	if result < n {
		return result
	}

	return -1
}

func abs(num int) int {
	if num < 0 {
		return -num
	}
	return num
}

func main() {
	// result: 1
	// words := []string{"hello","i","am","leetcode","hello"}
	// target := "hello"
	// startIndex := int(1)

	// result: 1
	// words := []string{"a","b","leetcode"}
	// target := "leetcode"
	// startIndex := int(0)

	// result: -1
	words := []string{"i","eat","leetcode"}
	target := "ate"
	startIndex := int(0)

	// result: 
	// words := []string{}
	// target := ""
	// startIndex := int(0)

	result := closestTarget(words, target, startIndex)
	fmt.Printf("result = %v\n", result)
}

