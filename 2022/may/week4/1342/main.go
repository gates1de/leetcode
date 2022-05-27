package main
import (
	"fmt"
)

func numberOfSteps(num int) int {
	result := int(0)
	for num > 0 {
		if num % 2 == 0 {
			num /= 2
		} else {
			num--
		}
		result++
	}

	return result
}

func main() {
	// result: 6
	// num := int(14)

	// result: 4
	// num := int(8)

	// result: 12
	// num := int(123)

	// result: 0
	num := int(0)

	// result: 
	// num := int(0)

	result := numberOfSteps(num)
	fmt.Printf("result = %v\n", result)
}

