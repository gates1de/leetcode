package main
import (
	"fmt"
)

func countSymmetricIntegers(low int, high int) int {
	result := int(0)

	for num := low; num <= high; num++ {
		if num < 100 && num % 11 == 0 {
			result++
		} else if 1000 <= num && num < 10000 {
			left := num / 1000 + (num % 1000) / 100
			right := (num % 100) / 10 + num % 10

			if left == right {
				result++
			}
		}
	}

	return result
}

func main() {
	// result: 9
	// low := int(1)
	// high := int(100)

	// result: 4
	low := int(1200)
	high := int(1230)

	// result: 
	// low := int(0)
	// high := int(0)

	result := countSymmetricIntegers(low, high)
	fmt.Printf("result = %v\n", result)
}

