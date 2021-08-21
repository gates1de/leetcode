package main
import (
	"fmt"
	"strconv"
)

func numDecodings(s string) int {
	if len(s) == 0 || s[0] == '0' {
		return 0
	} else if len(s) == 1 {
		return 1
	}

	result := int(1)
	prevCount := int(1)
	for i := 1; i < len(s); i++ {
		currCount := 0
		if s[i - 1] == '1' || (s[i - 1] == '2' && s[i] < '7') {
			currCount += prevCount
		}

		if s[i] > '0' {
			currCount += result
		}

		prevCount = result
		result = currCount
	}

	return result
}

// Time Limit Exceeded
func ngSolution(s string) int {
	if len(s) == 1 {
		if s == "0" {
			return 0
		}
		return 1
	}

	result := int(0)
	helper(1, s, &result)
	helper(2, s, &result)
	return result
}

func helper(digit int, s string, result *int) {
	// fmt.Printf("digit = %v, s = %v\n", digit, s)
	if len(s) == 0 {
		*result++
		return
	}

	char := s[:digit]
	// fmt.Printf("char = %v\n", char)
	num, _ := strconv.Atoi(char)
	if char[0] == byte('0') || (num <= 0 || 27 <= num) {
		return
	}

	helper(1, s[digit:], result)
	if len(s) >= 2 && len(s[digit:]) >= 2 {
		helper(2, s[digit:], result)
	}
}

func main() {
	// result: 2
	// s := "12"

	// result: 3
	// s := "226"

	// result: 0
	// s := "0"

	// result: 0
	// s := "06"

	// result: 
	s := "111111111111111111111111111111111111111111111"

	// result: 
	// s := ""

	result := numDecodings(s)
	fmt.Printf("result = %v\n", result)
}

