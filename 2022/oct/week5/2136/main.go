package main
import (
	"fmt"
	"sort"
)

func earliestFullBloom(plantTime []int, growTime []int) int {
	n := len(plantTime)
	indices := make([]int, n)
	for i, _ := range indices {
		indices[i] = i
	}

	sort.Slice(indices, func(a, b int) bool {
		return growTime[indices[a]] > growTime[indices[b]]
	})

	result := int(0)
	currentPlantTime := int(0)
	for _, index := range indices {
		time := currentPlantTime + plantTime[index] + growTime[index]
		result = max(result, time)
		currentPlantTime += plantTime[index]
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
	// result: 9
	// plantTime := []int{1,4,3}
	// growTime := []int{2,3,1}

	// result: 9
	// plantTime := []int{1,2,3,2}
	// growTime := []int{2,1,2,1}

	// result: 2
	plantTime := []int{1}
	growTime := []int{1}

	// result: 
	// plantTime := []int{}
	// growTime := []int{}

	result := earliestFullBloom(plantTime, growTime)
	fmt.Printf("result = %v\n", result)
}

