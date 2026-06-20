package main

import (
	"fmt"
)

func processStr(s string) string {
	result := make([]byte, 0, len(s))
	for i := 0; i < len(s); i++ {
		c := s[i]
		switch c {
		case '*':
			if len(result) > 0 {
				result = result[:len(result)-1]
			}
		case '#':
			result = append(result, result...)
		case '%':
			for l, r := 0, len(result)-1; l < r; l, r = l+1, r-1 {
				result[l], result[r] = result[r], result[l]
			}
		default:
			result = append(result, c)
		}
	}

	return string(result)
}

func main() {
	// result: "ba"
	// s := "a#b%*"

	// result: ""
	s := "z*#"

	// result: ""
	// s := ""

	result := processStr(s)
	fmt.Printf("result = %v\n", result)
}

