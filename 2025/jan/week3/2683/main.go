package main
import (
	"fmt"
)

func doesValidArrayExist(derived []int) bool {
	sum := int(0)
	for _, num := range derived {
		sum += num
	}

	return sum % 2 == 0
}

func main() {
	// result: true
	// derived := []int{1,1,0}

	// result: true
	// derived := []int{1,1}

	// result: false
	derived := []int{1,0}

	// result: 
	// derived := []int{}

	result := doesValidArrayExist(derived)
	fmt.Printf("result = %v\n", result)
}

