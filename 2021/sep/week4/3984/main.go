package main
import (
	"fmt"
)

func maxLength(arr []string) int {
	result := int(0)
	m := map[rune]bool{}
	helper(m, arr, &result)
	return result
}

func helper(m map[rune]bool, arr []string, result *int) {
	if len(arr) == 0 {
		// fmt.Printf("m = %v\n", m)
		if *result < len(m) {
			*result = len(m)
		}
		return
	}

	newM := copyMap(m)
	str := arr[0]
	for _, r := range str {
		if newM[r] {
			helper(m, arr[1:], result)
			return
		}
		newM[r] = true
	}
	helper(newM, arr[1:], result)
	helper(m, arr[1:], result)
}

func copyMap(m map[rune]bool) map[rune]bool {
	result := map[rune]bool{}
	for k, v := range m {
		result[k] = v
	}
	return result
}

func main() {
	// result: 4
	// arr := []string{"un", "iq", "ue"}

	// result: 6
	// arr := []string{"cha", "r", "act", "ers"}

	// result: 26
	// arr := []string{"abcdefghijklmnopqrstuvwxyz"}

	// result: 0
	arr := []string{"abcdefghijklmnopqrstuvwxya"}

	// result: 
	// arr := []string{}

	result := maxLength(arr)
	fmt.Printf("result = %v\n", result)
}

