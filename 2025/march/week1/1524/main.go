package main
import (
	"fmt"
)

const modulo = int(1e9 + 7)

func numOfSubarrays(arr []int) int {
	result := int(0)
	prefixSum := int(0)
	oddCount := int(0)
	evenCount := int(1)

	for _, num := range arr {
		prefixSum += num

		if prefixSum % 2 == 0 {
			result += oddCount
			evenCount++
		} else {
			result += evenCount
			oddCount++
		}

		result %= modulo
	}

	return result
}

func main() {
	// result: 4
	// arr := []int{1,3,5}

	// result: 0
	// arr := []int{2,4,6}

	// result: 16
	arr := []int{1,2,3,4,5,6,7}

	// result: 
	// arr := []int{}

	result := numOfSubarrays(arr)
	fmt.Printf("result = %v\n", result)
}

