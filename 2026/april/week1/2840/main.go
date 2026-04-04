package main

import (
	"fmt"
)

func checkStrings(s1 string, s2 string) bool {
	if len(s1) != len(s2) {
		return false
	}

	counts := make([]int, 256)
	for i := range len(s1) {
		offset := (i & 1) << 7
		counts[offset+int(s1[i])]++
		counts[offset+int(s2[i])]--
	}

	for _, count := range counts {
		if count != 0 {
			return false
		}
	}

	return true
}

func main() {
	// result: true
	// s1 := "abcd"
	// s2 := "cdab"

	// result: false
	s1 := "abcd"
	s2 := "dacb"

	// result: false
	// s1 := ""
	// s2 := ""

	result := checkStrings(s1, s2)
	fmt.Printf("result = %v\n", result)
}

