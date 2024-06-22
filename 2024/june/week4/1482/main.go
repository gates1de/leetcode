package main
import (
	"fmt"
)

func minDays(bloomDay []int, m int, k int) int {
	start := int(0)
	end := int(0)

	for _, day := range bloomDay {
		end = max(end, day)
	}

	result := int(-1)
	for start <= end {
		mid := (start + end) / 2

		if getNumOfBouquets(bloomDay, mid, k) >= m {
			result = mid
			end = mid - 1
		} else {
			start = mid + 1
		}
	}

	return result
}

func getNumOfBouquets(bloomDay []int, mid int, k int) int {
	result := int(0)
	count := int(0)

	for _, day := range bloomDay {
		if day <= mid {
			count++
		} else {
			count = 0
		}

		if count == k {
			result++
			count = 0
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
	// result: 3
	// bloomDay := []int{1,10,3,10,2}
	// m := int(3)
	// k := int(1)

	// result: -1
	// bloomDay := []int{1,10,3,10,2}
	// m := int(3)
	// k := int(2)

	// result: 12
	bloomDay := []int{7,7,7,7,12,7,7}
	m := int(2)
	k := int(3)

	// result: 
	// bloomDay := []int{}
	// m := int(0)
	// k := int(0)

	result := minDays(bloomDay, m, k)
	fmt.Printf("result = %v\n", result)
}

