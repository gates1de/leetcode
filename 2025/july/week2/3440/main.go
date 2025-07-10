package main

import (
	"fmt"
)

func maxFreeTime(eventTime int, startTime []int, endTime []int) int {
	n := len(startTime)
	q := make([]bool, n)
	t1 := int(0)
	t2 := int(0)

	for i := range startTime {
		if endTime[i]-startTime[i] <= t1 {
			q[i] = true
		}

		if i == 0 {
			t1 = max(t1, startTime[i])
		} else {
			t1 = max(t1, startTime[i] - endTime[i - 1])
		}

		if endTime[n - i - 1] - startTime[n - i - 1] <= t2 {
			q[n - i - 1] = true
		}

		if i == 0 {
			t2 = max(t2, eventTime - endTime[n - 1])
		} else {
			t2 = max(t2, startTime[n - i]- endTime[n - i - 1])
		}
	}

	result := int(0)
	for i := range startTime {
		left := int(0)
		if i != 0 {
			left = endTime[i - 1]
		}

		right := eventTime
		if i != n - 1 {
			right = startTime[i + 1]
		}

		if q[i] {
			result = max(result, right - left)
		} else {
			result = max(result, right - left - (endTime[i] - startTime[i]))
		}
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
	// result: 2
	// eventTime := int(5)
	// startTime := []int{1,3}
	// endTime := []int{2,5}

	// result: 7
	// eventTime := int(10)
	// startTime := []int{0,7,9}
	// endTime := []int{1,8,10}

	// result: 6
	// eventTime := int(10)
	// startTime := []int{0,3,7,9}
	// endTime := []int{1,4,8,10}

	// result: 0
	eventTime := int(5)
	startTime := []int{0,1,2,3,4}
	endTime := []int{1,2,3,4,5}

	// result: 
	// eventTime := int(0)
	// startTime := []int{}
	// endTime := []int{}

	result := maxFreeTime(eventTime, startTime, endTime)
	fmt.Printf("result = %v\n", result)
}

