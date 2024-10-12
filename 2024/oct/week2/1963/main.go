package main
import (
	"fmt"
)

func minSwaps(s string) int {
	stackSize := int(0)

	for _, char := range s {
		if char == '[' {
			stackSize++
		} else {
			if stackSize > 0 {
				stackSize--
			}
		}

	}

	return (stackSize + 1) / 2
}

func main() {
	// result: 1
	// s := "][]["

	// result: 2
	// s := "]]][[["

	// result: 0
	s := "[]"

	// result: 
	// s := ""

	result := minSwaps(s)
	fmt.Printf("result = %v\n", result)
}

