package main
import (
	"fmt"
	"sort"
)

func eliminateMaximum(dist []int, speed []int) int {
	arrival := make([]float64, len(dist))
	for i, d := range dist {
		arrival[i] = float64(d) / float64(speed[i])
	}

	sort.Slice(arrival, func(a, b int) bool { return arrival[a] < arrival[b] })
	result := int(0)
	for i, a := range arrival {
		if a <= float64(i) {
			break
		}

		result++
	}

	return result
}

func main() {
	// result: 3
	// dist := []int{1,3,4}
	// speed := []int{1,1,1}

	// result: 1
	// dist := []int{1,1,2,3}
	// speed := []int{1,1,1,1}

	// result: 1
	dist := []int{3,2,4}
	speed := []int{5,3,2}

	// result: 
	// dist := []int{}
	// speed := []int{}

	result := eliminateMaximum(dist, speed)
	fmt.Printf("result = %v\n", result)
}

