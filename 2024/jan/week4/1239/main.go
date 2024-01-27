package main
import (
	"fmt"
)

func maxLength(arr []string) int {
	if len(arr) == 0 {
		return 0
	}

	result := int(0)
	dfs(arr, "", 0, &result)
	return result
}

func dfs(arr []string, path string, i int, result *int) {
	isUniqueChar := checkUniqueChars(path)

	if isUniqueChar {
		*result = max(*result, len(path))
	}

	if i == len(arr) || !isUniqueChar {
		return
	}

	for j := i; j < len(arr); j++ {
		dfs(arr, path + arr[j], j + 1, result)
	}
}

func checkUniqueChars(s string) bool {
	m := make(map[rune]bool)
	for _, char := range s {
		if m[char] {
			return false
		}
		m[char] = true
	}
	return true
}

func max(a, b int) int {
	if b > a {
		return b
	}
	return a
}

func main() {
	// result: 4
	// arr := []string{"un","iq","ue"}

	// result: 6
	// arr := []string{"cha","r","act","ers"}

	// result: 26
	// arr := []string{"abcdefghijklmnopqrstuvwxyz"}

	// result: 6
	arr := []string{"a","abc","d","de","def"}

	// result: 
	// arr := []string{}

	result := maxLength(arr)
	fmt.Printf("result = %v\n", result)
}

