package main

import (
	"fmt"
)

const modulo = int(1e9 + 7)

func possibleStringCount(word string, k int) int {
	n := len(word)
	count := int(1)
	freq := make([]int, 0, n)

	for i := 1; i < n; i++ {
		if word[i] == word[i - 1] {
			count++
		} else {
			freq = append(freq, count)
			count = 1
		}
	}
	freq = append(freq, count)

	result := int(1)
	for _, c := range freq {
		result = result * c % modulo
	}

	if len(freq) >= k {
		return result
	}

	f := make([]int, k)
	g := make([]int, k)
	f[0] = 1
	for i := range g {
		g[i] = 1
	}

	for i := 0; i < len(freq); i++ {
		newF := make([]int, k)
		for j := 1; j < k; j++ {
			newF[j] = g[j - 1]

			if j - freq[i]-1 >= 0 {
				newF[j] = (newF[j] - g[j - freq[i] - 1] + modulo) % modulo
			}
		}

		newG := make([]int, k)
		newG[0] = newF[0]
		for j := 1; j < k; j++ {
			newG[j] = (newG[j - 1] + newF[j]) % modulo
		}

		f, g = newF, newG
	}

	return (result - g[k - 1] + modulo) % modulo
}

func main() {
	// result: 5
	// word := "aabbccdd"
	// k := int(7)

	// result: 1
	// word := "aabbccdd"
	// k := int(8)

	// result: 8
	word := "aaabbb"
	k := int(3)

	// result: 
	// word := ""
	// k := int(0)

	result := possibleStringCount(word, k)
	fmt.Printf("result = %v\n", result)
}

