package main
import (
	"fmt"
	"math"
)

func minSpeedOnTime(dist []int, hour float64) int {
	left := int(1)
	right := int(10000000)
	result := int(-1)

	for left <= right {
		mid := (left + right) / 2

		if timeRequired(dist, mid) <= hour {
			result = mid
			right = mid - 1
		} else {
			left = mid + 1
		}
	}

	return result
}

func timeRequired(dist []int, speed int) float64 {
	result := float64(0)

	for i, d := range dist {
		t := float64(d) / float64(speed)
		if i == len(dist) - 1 {
			result += t
		} else {
			result += math.Ceil(t)
		}
	}

	return result
}

func main() {
	// result: 1
	// dist := []int{1,3,2}
	// hour := float64(6)

	// result: 3
	// dist := []int{1,3,2}
	// hour := float64(2.7)

	// result: -1
	dist := []int{1,3,2}
	hour := float64(-1.9)

	// result: 
	// dist := []int{}
	// hour := float64(0)

	result := minSpeedOnTime(dist, hour)
	fmt.Printf("result = %v\n", result)
}

