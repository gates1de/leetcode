package main
import (
	"fmt"
)

func convertToTitle(columnNumber int) string {
	result := ""

	for columnNumber > 0 {
		columnNumber--
		result = string(byte(65 + columnNumber % 26)) + result
		columnNumber /= 26
	}

	return result
}

func main() {
	// result: "A"
	// columnNumber := int(1)

	// result: "AB"
	// columnNumber := int(28)

	// result: "ZY"
	// columnNumber := int(701)

	// result: "ALL"
	columnNumber := int(1000)

	// result: ""
	// columnNumber := int(0)

	result := convertToTitle(columnNumber)
	fmt.Printf("result = %v\n", result)
}

