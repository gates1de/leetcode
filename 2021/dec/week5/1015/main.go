package main
import (
	"fmt"
)

func smallestRepunitDivByK(k int) int {
	if k % 2 == 0 || k % 5 == 0 {
		return -1
	}

	current := int(0)
	for i := 1; i <= k; i++ {
		current = (current * 10 + 1) % k
		if current == 0 {
			return i
		}
	}

	return -1
}

func main() {
	// result: 1
	// k := int(1)

	// result: -1
	// k := int(2)

	// result: 3
	// k := int(3)

	// result: -1
	// k := int(4)

	// result: -1
	// k := int(5)

	// result: 6
	// k := int(7)

	// result: 45
	k := int(99999)

	// result: 
	// k := int(0)

	result := smallestRepunitDivByK(k)
	fmt.Printf("result = %v\n", result)
}

