package main
import (
	"fmt"
	"math"
	"sort"
)

func makeArrayIncreasing(arr1 []int, arr2 []int) int {
	dp := make(map[int]int)
	sort.Ints(arr2)
	n := len(arr2)
	dp[-1] = 0

	for i := 0; i < len(arr1); i++ {
		newDp := make(map[int]int)

		for prev, _ := range dp {
			key := arr1[i]
			if key > prev {
				val := math.MaxInt32
				if _, ok := newDp[key]; ok {
					val = newDp[key]
				}
				val = min(val, dp[prev])
				newDp[key] = val
			}

			index := bisectRight(arr2, prev)
			if index < n {
				val := math.MaxInt32
				key := arr2[index]
				if _, ok := newDp[key]; ok {
					val = newDp[key]
				}
				val = min(val, 1 + dp[prev])
				newDp[key] = val
			}
		}

		dp = newDp
	}

	result := math.MaxInt32
	for _, val := range dp {
		result = min(result, val)
	}

	if result == math.MaxInt32 {
		return -1
	}
	return result
}

func bisectRight(arr []int, value int) int {
	left := int(0)
	right := len(arr)
	for left < right {
		mid := (left + right) / 2
		if arr[mid] <= value {
			left = mid + 1
		} else {
			right = mid
		}
	}
	return left
}

func min(a, b int) int {
	if b < a {
		return b
	}
	return a
}

func main() {
	// result: 1
	// arr1 := []int{1,5,3,6,7}
	// arr2 := []int{1,3,2,4}

	// result: 2
	// arr1 := []int{1,5,3,6,7}
	// arr2 := []int{4,3,1}

	// result: -1
	arr1 := []int{1,5,3,6,7}
	arr2 := []int{1,6,3,3}

	// result: 
	// arr1 := []int{}
	// arr2 := []int{}

	result := makeArrayIncreasing(arr1, arr2)
	fmt.Printf("result = %v\n", result)
}

