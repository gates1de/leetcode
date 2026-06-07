package main

import (
	"fmt"
)

func totalWaviness(num1 int, num2 int) int {
	result := int(0)
	for num := num1; num <= num2; num++ {
		result += waviness(num)
	}
	return result
}

func waviness(num int) int {
	digits := make([]int, 0)
	for num > 0 {
		digits = append(digits, num % 10)
		num /= 10
	}

	result := int(0)
	for left := len(digits) - 2; left >= 1; left-- {
		mid := digits[left]
		if mid > digits[left-1] && mid > digits[left+1] {
			result++
		}
		if mid < digits[left-1] && mid < digits[left+1] {
			result++
		}
	}

	return result
}

func main() {
	// result: 3
	// num1 := int(120)
	// num2 := int(130)

	// result: 3
	// num1 := int(198)
	// num2 := int(202)

	// result: 2
	num1 := int(4848)
	num2 := int(4848)

	// result: 
	// num1 := int(0)
	// num2 := int(0)

	result := totalWaviness(num1, num2)
	fmt.Printf("result = %v\n", result)
}
