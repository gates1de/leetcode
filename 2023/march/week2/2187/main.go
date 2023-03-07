package main
import (
	"fmt"
)

func minimumTime(time []int, totalTrips int) int64 {
	maxTime := int(0)
	for _, t := range time {
		maxTime = max(maxTime, t)
	}
	left := int64(1)
	right := int64(maxTime * totalTrips)

	for left < right {
		mid := (left + right) / 2
		if timeEnough(time, mid, totalTrips) {
			right = mid
		} else {
			left = mid + 1
		}
	}

	return left
}

func timeEnough(time []int, givenTime int64, totalTrips int) bool {
	actualTrips := int64(0)
	for _, t := range time {
		actualTrips += givenTime / int64(t)
	}
	return actualTrips >= int64(totalTrips)
}

func max(a, b int) int {
	if b > a {
		return b
	}
	return a
}

func main() {
	// result: 3
	// time := []int{1,2,3}
	// totalTrips := int(5)

	// result: 2
	time := []int{2}
	totalTrips := int(1)

	// result: 
	// time := []int{}
	// totalTrips := int(0)

	result := minimumTime(time, totalTrips)
	fmt.Printf("result = %v\n", result)
}

