package main
import (
	"fmt"
)

func numberOfSteps(num int) int {
	result := int(0)
	for num > 0 {
		if num % 2 == 0 {
			num = num / 2
		} else {
			num--
		}
		result++
	}
	return result
}

func main() {
	// num := int(14)
	// num := int(8)
	num := int(123)
	result := numberOfSteps(num)
	fmt.Printf("result = %v\n", result)
}

