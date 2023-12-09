package main
import (
	"fmt"
)

func largestOddNumber(num string) string {
	for i := len(num) - 1; i >= 0; i-- {
		char := num[i]
		if char == '1' || char == '3' || char == '5' || char == '7' || char == '9' {
			return num[:i + 1]
		}
	}

	return ""
}

func main() {
	// result: "5"
	// num := "52"

	// result: ""
	// num := "4206"

	// result: "35427
	num := "35427"

	// result: ""
	// num := ""

	result := largestOddNumber(num)
	fmt.Printf("result = %v\n", result)
}

