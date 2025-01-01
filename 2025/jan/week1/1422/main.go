package main
import (
	"fmt"
)

func maxScore(s string) int {
	if len(s) <= 1 {
		return 0
	} else if len(s) == 2 {
		if s == "00" || s == "11" {
			return 1
		} else if s == "01" {
			return 2
		}
		return 0
	}

	zeroCounts := make(map[int]int)
	oneCounts := make(map[int]int)
	zeroCount := int(0)
	oneCount := int(0)

	for i, _ := range s {
		j := len(s) - i - 1
		left := s[i]
		right := s[j]

		if left == '0' {
			zeroCount++
		}
		if right == '1' {
			oneCount++
		}

		zeroCounts[i] = zeroCount
		oneCounts[j] = oneCount
	}

	result := int(0)
	for i := 1; i < len(s) - 1; i++ {
		result = max(result, zeroCounts[i] + oneCounts[i])
	}

	return result
}

func max(a, b int) int {
	if b > a {
		return b
	}
	return a
}

func main() {
	// result: 5
	// s := "011101"

	// result: 5
	// s := "00111"

	// result: 3
	s := "1111"

	// result: 
	// s := ""

	result := maxScore(s)
	fmt.Printf("result = %v\n", result)
}
