package main
import (
	"fmt"
)

func threeConsecutiveOdds(arr []int) bool {
	count := int(0)
	for _, num := range arr {
		if num % 2 != 0 {
			count++
		} else {
			count = 0
		}

		if count >= 3 {
			return true
		}
	}

	return false
}

func main() {
	// result: false
	// arr := []int{2,6,4,1}

	// result: true
	arr := []int{1,2,34,3,4,5,7,23,12}

	// result: 
	// arr := []int{}

	result := threeConsecutiveOdds(arr)
	fmt.Printf("result = %v\n", result)
}
