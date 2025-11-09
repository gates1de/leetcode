package main

import (
	"fmt"
)

func countOperations(num1 int, num2 int) int {
	return operation(0, num1, num2)
}

func operation(count int, num1 int, num2 int) int {
	if num1 == 0 || num2 == 0 {
		return count
	} else if num1 < num2 {
		return operation(count + 1, num1, num2 - num1)
	}
	return operation(count + 1, num1 - num2, num2)
}

func main() {
	// result: 3
	// num1 := int(2)
	// num2 := int(3)

	// result: 1
	num1 := int(10)
	num2 := int(10)

	// result: 
	// num1 := int(0)
	// num2 := int(0)

	result := countOperations(num1, num2)
	fmt.Printf("result = %v\n", result)
}

