package main
import (
	"fmt"
)

func garbageCollection(garbage []string, travel []int) int {
	for i := 1; i < len(travel); i++ {
		travel[i] += travel[i - 1]
	}

	lastPositions := make(map[byte]int)
	result := int(0)

	for i, g := range garbage {
		for _, char := range []byte(g) {
			lastPositions[char] = i
		}
		result += len(g)
	}

	garbageTypes := "MPG"
	for _, char := range []byte(garbageTypes) {
		if lastPositions[char] > 0 {
			result += travel[lastPositions[char] - 1]
		}
	}

	return result
}

func main() {
	// result: 21
	// garbage := []string{"G","P","GP","GG"}
	// travel := []int{2,4,3}

	// result: 37
	garbage := []string{"MMM","PGM","GP"}
	travel := []int{3,10}

	// result: 
	// garbage := []string{}
	// travel := []int{}

	result := garbageCollection(garbage, travel)
	fmt.Printf("result = %v\n", result)
}

