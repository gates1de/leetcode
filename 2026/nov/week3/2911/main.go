package main

import (
	"fmt"
)

func minimumChanges(s string, k int) int {
	n := len(s)
	changes := make([]int, n*n)
	for length := 2; length <= n; length++ {
		for start := 0; start+length <= n; start++ {
			best := n
			for divisor := 1; divisor < length; divisor++ {
				if length%divisor != 0 {
					continue
				}

				current := 0
				for offset := 0; offset < divisor; offset++ {
					left := start + offset
					right := left + length - divisor
					for left < right {
						if s[left] != s[right] {
							current++
						}
						left += divisor
						right -= divisor
					}
				}
				best = min(best, current)
			}
			changes[start*n+start+length-1] = best
		}
	}

	const inf = int(1e9)
	previous := make([]int, n+1)
	current := make([]int, n+1)
	for i := range previous {
		previous[i] = inf
	}
	previous[0] = 0

	for parts := 1; parts <= k; parts++ {
		for i := range current {
			current[i] = inf
		}

		for end := parts * 2; end <= n; end++ {
			for start := (parts - 1) * 2; start <= end-2; start++ {
				candidate := previous[start] + changes[start*n+end-1]
				current[end] = min(current[end], candidate)
			}
		}

		previous, current = current, previous
	}

	return previous[n]
}

func main() {
	// result: 1
	// s := "abcac"
	// k := 2

	// result: 2
	// s := "abcdef"
	// k := 2

	// result: 0
	s := "aabbaa"
	k := 3

	result := minimumChanges(s, k)
	fmt.Printf("result = %v\n", result)
}
