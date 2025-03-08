package main
import (
	"fmt"
)

func coloredCells(n int) int64 {
	return 1 + int64(n * (n - 1) * 2)
}

func main() {
	// result: 1
	// n := int(1)

	// result: 5
	n := int(2)

	// result: 
	// n := int(0)

	result := coloredCells(n)
	fmt.Printf("result = %v\n", result)
}

