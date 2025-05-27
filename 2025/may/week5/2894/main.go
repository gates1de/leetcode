package main
import (
	"fmt"
)

func differenceOfSums(n int, m int) int {
	num1 := int(0)
	num2 := int(0)

	for num := 1; num <= n; num++ {
		if num % m != 0 {
			num1 += num
		} else {
			num2 += num
		}
	}

	return num1 - num2
}

func main() {
	// result: 19
	// n := int(10)
	// m := int(3)

	// result: 15
	// n := int(5)
	// m := int(6)

	// result: -15
	n := int(5)
	m := int(1)

	// result: 
	// n := int(0)
	// m := int(0)

	result := differenceOfSums(n, m)
	fmt.Printf("result = %v\n", result)
}

