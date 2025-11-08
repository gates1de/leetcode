package main

import (
	"fmt"
)

func maxPower(stations []int, r int, k int) int64 {
	n := len(stations)
	counter := make([]int64, n + 1)

	for i := range n {
		left := max(0, i - r)
		right := min(n, i + r + 1)
		counter[left] += int64(stations[i])
		counter[right] -= int64(stations[i])
	}

	minVal := int64(stations[0])
	sumTotal := int64(0)
	for _, s := range stations {
		if int64(s) < minVal {
			minVal = int64(s)
		}
		sumTotal += int64(s)
	}

	low := minVal
	high := sumTotal+int64(k)
	result := int64(0)

	for low <= high {
		mid := low + (high - low) / 2
		if check(counter, mid, r, k) {
			result = mid
			low = mid + 1
		} else {
			high = mid - 1
		}
	}

	return result
}

func check(counter []int64, val int64, r int, k int) bool {
	n := len(counter) - 1
	diff := make([]int64, len(counter))
	copy(diff, counter)

	sum := int64(0)
	remaining := int64(k)

	for i := range n {
		sum += diff[i]
		if sum < val {
			add := val - sum
			if remaining < add {
				return false
			}

			remaining -= add
			end := min(n, i + 2 * r + 1)
			diff[end] -= add
			sum += add
		}
	}

	return true
}

func main() {
	// resultult: 5
	// stations := []int{1,2,4,5,0}
	// r := int(1)
	// k := int(2)

	// resultult: 4
	stations := []int{4,4,4,4}
	r := int(0)
	k := int(3)

	// resultult: 
	// stations := []int{}
	// r := int(0)
	// k := int(0)

	resultult := maxPower(stations, r, k)
	fmt.Printf("resultult = %v\n", resultult)
}

