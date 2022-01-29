package main
import (
	"fmt"
)

func sequentialDigits(low int, high int) []int {
	minDigit := int(0)
	maxDigit := int(0)
	for i := 1; i <= low; i *= 10 {
		minDigit++
	}
	for i := 1; i <= high; i *= 10 {
		maxDigit++
	}

	result := []int{}
	for digit := minDigit; digit <= maxDigit; digit++ {
		for j := 1; j <= 10 - digit; j++ {
			num := int(0)
			for k := j; k < j + digit; k++ {
				num += k
				num *= 10
			}
			num /= 10
			if num < low || high < num {
				continue
			}
			result = append(result, num)
			// fmt.Println(num)
		}
	}

	return result
}

func main() {
	// result: [123,234]
	// low := int(100)
	// high := int(300)

	// result: [1234,2345,3456,4567,5678,6789,12345]
	// low := int(1000)
	// high := int(13000)

	// result: [12,23,34,45,56,67,78,89,123,234,345,456,567,678,789,1234,2345,3456,4567,5678,6789,12345,23456,34567,45678,56789,123456,234567,345678,456789,1234567,2345678,3456789,12345678,23456789,123456789]
	// low := int(10)
	// high := int(1000000000)

	// result: [23,34,45,56,67,78,89,123,234,345,456,567,678,789,1234,2345,3456,4567,5678,6789,12345,23456,34567,45678,56789,123456,234567,345678,456789,1234567,2345678,3456789,12345678,23456789,123456789]
	low := int(13)
	high := int(999999999)

	// result: 
	// low := int(0)
	// high := int(0)

	result := sequentialDigits(low, high)
	fmt.Printf("result = %v\n", result)
}

