package main

import (
	"fmt"
	"strings"
)

func checkOnesSegment(s string) bool {
	return !strings.Contains(s, "01")
}

func main() {
	// result: false
	// s := "1001"

	// result: true
	s := "110"

	// result: false
	// s := ""

	result := checkOnesSegment(s)
	fmt.Printf("result = %v\n", result)
}

