package main

import (
	"fmt"
	"math"
)

const eps = float64(1e-7)

func minNumberOfSeconds(mountainHeight int, workerTimes []int) int64 {
	maxWorkerTimes := int(0)
	for _, t := range workerTimes {
		maxWorkerTimes = max(maxWorkerTimes, t)
	}

	l := int64(1)
	r := int64(maxWorkerTimes) * int64(mountainHeight) * int64(mountainHeight + 1) / 2
	result := int64(0)

	for l <= r {
		mid := (l + r) / 2
		count := int64(0)

		for _, t := range workerTimes {
			work := mid / int64(t)
			k := int64((-1.0 + math.Sqrt(1 + float64(work) * 8)) / 2 + eps)
			count += k
		}

		if count >= int64(mountainHeight) {
			result = mid
			r = mid - 1
		} else {
			l = mid + 1
		}
	}

	return result
}

func main() {
	// result: 3
	// mountainHeight := int(4)
	// workerTimes := []int{2,1,1}

	// result: 12
	// mountainHeight := int(10)
	// workerTimes := []int{3,2,2,4}

	// result: 15
	mountainHeight := int(5)
	workerTimes := []int{1}

	// result: 
	// mountainHeight := int(0)
	// workerTimes := []int{}

	result := minNumberOfSeconds(mountainHeight, workerTimes)
	fmt.Printf("result = %v\n", result)
}

