package main
import (
	"fmt"
	"sort"
)

func maxSatisfaction(satisfaction []int) int {
	sort.Ints(satisfaction)
	result := int(0)
	suffixSum := int(0)

	for i := len(satisfaction) - 1; i >= 0 && suffixSum + satisfaction[i] > 0; i-- {
        suffixSum += satisfaction[i]
        result +=  suffixSum
	}
	return result
}

func main() {
	// result: 14
	// satisfaction := []int{-1,-8,0,5,-9}

	// result: 20
	// satisfaction := []int{4,3,2}

	// result: 0
	satisfaction := []int{-1,-4,-5}

	// result: 
	// satisfaction := []int{}

	result := maxSatisfaction(satisfaction)
	fmt.Printf("result = %v\n", result)
}

