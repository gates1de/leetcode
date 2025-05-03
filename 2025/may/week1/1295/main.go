package main
import (
	"fmt"
)

func findNumbers(nums []int) int {
	result := int(0)

	for _, num := range nums {
		if getDigit(num) % 2 == 0 {
			result++
		}
	}

	return result
}

func getDigit(num int) int {
	result := int(0)
	for n := num; n > 0; n /= 10 {
		result++
	}
	return result
}

func main() {
	// result: 2
	// nums := []int{12,345,2,6,7896}

	// result: 1
	nums := []int{555,901,482,1771}

	// result: 
	// nums := []int{}

	result := findNumbers(nums)
	fmt.Printf("result = %v\n", result)
}

