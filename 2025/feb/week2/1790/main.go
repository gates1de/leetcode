package main
import (
	"fmt"
)

func areAlmostEqual(s1 string, s2 string) bool {
	firstIndexDiff := int(0)
	secondIndexDiff := int(0)
	numDiffs := int(0)

	for i, _ := range s1 {
		if s1[i] != s2[i] {
			numDiffs++

			if numDiffs > 2 {
				return false
			} else if numDiffs == 1 {
				firstIndexDiff = i
			} else {
				secondIndexDiff = i
			}
		}
	}

	return s1[firstIndexDiff] == s2[secondIndexDiff] && s1[secondIndexDiff] == s2[firstIndexDiff]
}

func main() {
	// result: true
	// s1 := "bank"
	// s2 := "kanb"

	// result: false
	// s1 := "attack"
	// s2 := "defend"

	// result: true
	s1 := "kelb"
	s2 := "kelb"

	// result: 
	// s1 := ""
	// s2 := ""

	result := areAlmostEqual(s1, s2)
	fmt.Printf("result = %v\n", result)
}

