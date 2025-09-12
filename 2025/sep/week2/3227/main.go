package main

import (
	"fmt"
	"strings"
)

func doesAliceWin(s string) bool {
	return strings.ContainsAny(s, "aeiou")
}

func main() {
	// result: true
	// s := "leetcoder"

	// result: false
	s := "bbcd"

	// result: 
	// s := ""

	result := doesAliceWin(s)
	fmt.Printf("result = %v\n", result)
}

