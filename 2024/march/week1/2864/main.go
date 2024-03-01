package main
import (
	"fmt"
)

func maximumOddBinaryNumber(s string) string {
	n := len(s)
	if n == 1 {
		return s
	}

	results := make([]byte, 0, n)
	for _, char := range s {
		results = append(results, '0')
		if char == '1' {
			copy(results[1:], results)
			results[0] = '1'
		}
	}

	if results[n - 1] == '0' {
		copy(results, results[1:])
		results[n - 1] = '1'
	}
    return string(results)
}

func main() {
	// result: "001"
	// s := "010"

	// result: "1001"
	s := "0101"

	// result: 
	// s := ""

	result := maximumOddBinaryNumber(s)
	fmt.Printf("result = %v\n", result)
}

