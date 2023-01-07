package main
import (
	"fmt"
	"sort"
)

func minDeletionSize(strs []string) int {
	if len(strs) == 0 {
		return 0
	}

	columnStrs := make([]string, len(strs[0]))
	for i := 0; i < len(columnStrs); i++ {
		for _, str := range strs {
			columnStrs[i] += string(str[i])
		}
	}

	result := int(0)
	for _, str := range columnStrs {
		if !isLexicographically(str) {
			result++
		}
	}
	return result
}

func isLexicographically(str string) bool {
	if len(str) <= 1 {
		return true
	}

	chars := []byte(str)
	sort.Slice(chars, func(a, b int) bool {
		return chars[a] < chars[b]
	})
	return string(chars) == str
}

func main() {
	// result: 1
	// strs := []string{"cba","daf","ghi"}

	// result: 0
	// strs := []string{"a","b"}

	// result: 3
	strs := []string{"zyx","wvu","tsr"}

	// result: 
	// strs := []string{}

	result := minDeletionSize(strs)
	fmt.Printf("result = %v\n", result)
}

