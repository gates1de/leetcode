package main
import (
	"fmt"
)

func largestAltitude(gain []int) int {
	result := int(0)
	alt := int(0)
	for _, num := range gain {
		alt += num
		result = max(result, alt)
	}
	return result
}

func max(a, b int) int {
	if b > a {
		return b
	}
	return a
}

func main() {
	// result: 1
	// gain := []int{-5,1,5,0,-7}

	// result: 0
	gain := []int{-4,-3,-2,-1,4,3,2}

	// result: 
	// gain := []int{}

	result := largestAltitude(gain)
	fmt.Printf("result = %v\n", result)
}

