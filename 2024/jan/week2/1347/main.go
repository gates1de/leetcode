package main
import (
	"fmt"
)

func minSteps(s string, t string) int {
	count := make([]int, 26)
	for i, _ := range s {
		count[s[i] - 'a']--
		count[t[i] - 'a']++
	}

	result := int(0)
	for _, c := range count {
		result += max(0, c)
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
	// result: 1
	// s := "bab"
	// t := "aba"

	// result: 5
	// s := "leetcode"
	// t := "practice"

	// result: 0
	s := "anagram"
	t := "mangaar"

	// result: 
	// s := ""
	// t := ""

	result := minSteps(s, t)
	fmt.Printf("result = %v\n", result)
}

