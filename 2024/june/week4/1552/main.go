package main
import (
	"fmt"
	"math"
	"sort"
)

func maxDistance(position []int, m int) int {
	result := int(0)
	n := len(position)
	sort.Ints(position)

	low := int(1)
	high := int(math.Ceil(float64(position[n - 1]) / float64(m - 1)))

	for low <= high {
		mid := low + (high - low) / 2
		if canPlaceBalls(mid, position, m) {
			result = mid
			low = mid + 1
		} else {
			high = mid - 1
		}
	}

	return result
}

func canPlaceBalls(x int, position []int, m int) bool {
	prev := position[0]
	placed := int(1)

	for i := 1; i < len(position) && placed < m; i++ {
		current := position[i]
		if current - prev >= x {
			placed += 1
			prev = current
		}
	}

	return placed == m
}

func main() {
	// result: 3
	// position := []int{1,2,3,4,7}
	// m := int(3)

	// result: 999999999
	position := []int{5,4,3,2,1,1000000000}
	m := int(2)

	// result: 
	// position := []int{}
	// m := int(0)

	result := maxDistance(position, m)
	fmt.Printf("result = %v\n", result)
}

