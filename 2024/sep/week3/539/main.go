package main
import (
	"fmt"
	"math"
	"strconv"
)

func findMinDifference(timePoints []string) int {
	m := make([]bool, 24 * 60)
	for _, time := range timePoints {
		hour, _ := strconv.Atoi(time[:2])
		minute, _ := strconv.Atoi(time[3:])

		minutes := hour * 60 + minute
		if m[minutes] {
			return 0
		}
		m[minutes] = true
	}

	prevIndex := math.MaxInt32
	firstIndex := math.MaxInt32
	lastIndex := math.MaxInt32
	result := math.MaxInt32

	for i := 0; i < 24 * 60; i++ {
		if m[i] {
			if prevIndex != math.MaxInt32 {
				result = min(result, i - prevIndex)
			}

			prevIndex = i

			if firstIndex == math.MaxInt32 {
				firstIndex = i
			}

			lastIndex = i
		}
	}

	return min(result, 24 * 60 - lastIndex + firstIndex)
}

func min(a, b int) int {
	if b < a {
		return b
	}
	return a
}

func main() {
	// result: 1
	// timePoints := []string{"23:59","00:00"}

	// result: 0
	timePoints := []string{"00:00","23:59","00:00"}

	// result: 
	// timePoints := []string{}

	result := findMinDifference(timePoints)
	fmt.Printf("result = %v\n", result)
}

