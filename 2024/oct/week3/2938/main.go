package main
import (
	"fmt"
)

func minimumSteps(s string) int64 {
	result := int64(0)
	blackCount := int(0)

	for _, char := range s {
		if char == '0' {
			result += int64(blackCount)
		} else {
			blackCount++
		}
	}

	return result
}

func main() {
	// result: 1
	// s := "101"

	// result: 2
	// s := "100"

	// result: 0
	s := "0111"

	// result: 
	// s := ""

	result := minimumSteps(s)
	fmt.Printf("result = %v\n", result)
}

