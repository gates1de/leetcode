package main
import (
	"fmt"
)

func countOdds(low int, high int) int {
	diff := high - low
	if diff == 0 {
		if low % 2 == 1 {
			return 1
		}
		return 0
	}

	result := diff / 2
	if low % 2 == 1 || high % 2 == 1{
		result += 1
	}
	return result
}

func main() {
	// result: 3
	// low := int(3)
	// high := int(7)

	// result: 1
	// low := int(8)
	// high := int(10)

	// result: 1
	// low := int(1)
	// high := int(1)

	// result: 0
	// low := int(2)
	// high := int(2)

	// result: 1
	low := int(0)
	high := int(1)

	// result: 
	// low := int(0)
	// high := int(0)

	result := countOdds(low, high)
	fmt.Printf("result = %v\n", result)
}

