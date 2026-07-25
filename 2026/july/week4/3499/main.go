package main

import (
	"fmt"
)

func maxActiveSectionsAfterTrade(s string) int {
	active := int(0)
	for i := 0; i < len(s); i++ {
		if s[i] == '1' {
			active++
		}
	}

	var lengths []int
	var ones []bool
	for i := 0; i < len(s); {
		j := i + 1
		for j < len(s) && s[j] == s[i] {
			j++
		}

		lengths = append(lengths, j-i)
		ones = append(ones, s[i] == '1')
		i = j
	}

	prefix := make([]int, len(lengths)+1)
	for i := range lengths {
		prefix[i+1] = prefix[i]
		if !ones[i] && lengths[i] > prefix[i+1] {
			prefix[i+1] = lengths[i]
		}
	}

	suffix := make([]int, len(lengths)+1)
	for i := len(lengths) - 1; i >= 0; i-- {
		suffix[i] = suffix[i+1]
		if !ones[i] && lengths[i] > suffix[i] {
			suffix[i] = lengths[i]
		}
	}

	bestGain := int(0)
	for i := 1; i+1 < len(lengths); i++ {
		if !ones[i] || ones[i-1] || ones[i+1] {
			continue
		}

		if gain := lengths[i-1] + lengths[i+1]; gain > bestGain {
			bestGain = gain
		}

		otherZeroes := max(suffix[i+2], prefix[i-1])
		if gain := otherZeroes - lengths[i]; gain > bestGain {
			bestGain = gain
		}
	}

	return active + bestGain
}

func main() {
	// result: 1
	// s := "01"

	// result: 4
	// s := "0100"

	// result: 7
	s := "1000100"

	// result: 4
	// s := "01010"

	// result: 4
	// s := "0100"

	// result: 3
	// s := "1101"

	// result:
	// s := ""

	result := maxActiveSectionsAfterTrade(s)
	fmt.Printf("result = %v\n", result)
}
