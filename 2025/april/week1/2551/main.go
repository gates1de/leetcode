package main
import (
	"fmt"
	"sort"
)

func putMarbles(weights []int, k int) int64 {
	n := len(weights)
	pairWeights := make([]int, n - 1)
	for i, _ := range pairWeights[:n - 1] {
		pairWeights[i] = weights[i] + weights[i + 1]
	}
	sort.Ints(pairWeights[:n - 1])
	result := int64(0)
	for i := 0; i < k - 1; i++ {
		result += int64(pairWeights[n - 2 - i] - pairWeights[i])
	}

	return result
}

func main() {
	// result: 4
	// weights := []int{1,3,5,1}
	// k := int(2)

	// result: 0
	weights := []int{1,3}
	k := int(2)

	// result: 
	// weights := []int{}
	// k := int(0)

	result := putMarbles(weights, k)
	fmt.Printf("result = %v\n", result)
}

