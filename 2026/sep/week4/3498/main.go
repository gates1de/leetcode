package main

import (
	"fmt"
)

func reverseDegree(s string) int {
	result := int(0)
	for i := range s {
		result += int('z'-s[i]+1) * (i + 1)
	}
	return result
}

func main() {
	// result: 148
	// s := "abc"

	// result: 160
	s := "zaza"

	// result: 0
	// s := ""

	result := reverseDegree(s)
	fmt.Printf("result = %v\n", result)
}
