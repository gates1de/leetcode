package main

import (
	"fmt"
	"sort"
	"strings"
)

func makeLargestSpecial(s string) string {
	count := int(0)
	i := int(0)
	result := make([]string, 0, len(s))

	for j, char := range s {
		if char == '1' {
			count++
		} else {
			count--
		}

		if count == 0 {
			result = append(result, "1" + makeLargestSpecial(s[i + 1:j]) + "0")
			i = j + 1
		}
	}

	sort.Sort(sort.Reverse(sort.StringSlice(result)))
	return strings.Join(result, "")
}

func main() {
	// result: "11100100"
	// s := "11011000"

	// result: "10"
	s := "10"

	// result: 
	// s := ""

	result := makeLargestSpecial(s)
	fmt.Printf("result = %v\n", result)
}

