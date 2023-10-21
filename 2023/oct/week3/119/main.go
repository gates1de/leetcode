package main
import (
	"fmt"
)

func getRow(rowIndex int) []int {
	result := make([]int, 1, rowIndex + 1)
	result[0] = 1
    prev := int(1)
    for i := 1; i <= rowIndex; i++ {
        next := prev * (rowIndex - i + 1) / i
        result = append(result, next)
        prev = next
    }

    return result
}

func main() {
	// result: [1,3,3,1]
	// rowIndex := int(3)

	// result: [1]
	// rowIndex := int(0)

	// result: [1,1]
	rowIndex := int(1)

	// result: []
	// rowIndex := int(0)

	result := getRow(rowIndex)
	fmt.Printf("result = %v\n", result)
}

