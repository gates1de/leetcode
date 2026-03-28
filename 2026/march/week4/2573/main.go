package main

import (
	"fmt"
)

func findTheString(lcp [][]int) string {
	n := len(lcp)
	word := make([]byte, n)
	current := byte('a')

	for i := range n {
		if word[i] == 0 {
			if current > 'z' {
				return ""
			}

			word[i] = current
			for j := i + 1; j < n; j++ {
				if lcp[i][j] > 0 {
					word[j] = word[i]
				}
			}

			current++
		}
	}

	for i := n - 1; i >= 0; i-- {
		for j := n - 1; j >= 0; j-- {
			if word[i] != word[j] {
				if lcp[i][j] != 0 {
					return ""
				}
			} else {
				if i == n - 1 || j == n - 1 {
					if lcp[i][j] != 1 {
						return ""
					}
				} else {
					if lcp[i][j] != lcp[i + 1][j + 1] + 1 {
						return ""
					}
				}
			}
		}
	}

	return string(word)
}

func main() {
	// result: "abab"
	// lcp := [][]int{{4,0,2,0},{0,3,0,1},{2,0,2,0},{0,1,0,1}}

	// result: "aaaa"
	// lcp := [][]int{{4,3,2,1},{3,3,2,1},{2,2,2,1},{1,1,1,1}}

	// result: ""
	lcp := [][]int{{4,3,2,1},{3,3,2,1},{2,2,2,1},{1,1,1,3}}

	// result: ""
	// lcp := [][]int{}

	result := findTheString(lcp)
	fmt.Printf("result = %v\n", result)
}

