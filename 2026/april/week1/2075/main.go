package main

import (
	"fmt"
)

func decodeCiphertext(encodedText string, rows int) string {
	n := len(encodedText)
	cols := n / rows
	result := make([]byte, 0, n)

	for i := range cols {
		for j := i; j < n; j += cols + 1 {
			result = append(result, encodedText[j])
		}
	}

	for len(result) > 0 && result[len(result) - 1] == ' ' {
		result = result[:len(result) - 1]
	}

	return string(result)
}

func main() {
	// result: "cipher"
	// encodedText := "ch   ie   pr"
	// rows := int(3)

	// result: "i love leetcode"
	// encodedText := "iveo    eed   l te   olc"
	// rows := int(4)

	// result: "coding"
	encodedText := "coding"
	rows := int(1)

	// result: ""
	// encodedText := ""
	// rows := int(0)

	result := decodeCiphertext(encodedText, rows)
	fmt.Printf("result = %v\n", result)
}

