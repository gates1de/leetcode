package main
import (
	"fmt"
	"sort"
)

func findClosestElements(arr []int, k int, x int) []int {
	n := len(arr)
	xIndex := sort.SearchInts(arr, x)
	if xIndex >= n {
		return arr[max(0, n - k):]
	}

	i := xIndex
	j := xIndex

	if i == 0 {
		return arr[:k]
	}
	if j >= n {
		return arr[n - k:]
	}

	if arr[xIndex] != x {
		leftDiff := abs(x, arr[i - 1])
		rightDiff := abs(x, arr[j])
		if leftDiff > rightDiff {
			j++
		}
	}

	for j - i < k {
		if i - 1 < 0 {
			j++
			continue
		}

		if j + 1 > n {
			i--
			continue
		}

		leftDiff := abs(x, arr[i - 1])
		rightDiff := abs(x, arr[j])
		if leftDiff <= rightDiff {
			i--
		} else {
			j++
		}
	}

	return arr[i:j]
}

func max(a, b int) int {
	if b > a {
		return b
	}
	return a
}

func abs(a, b int) int {
	if b > a {
		return b - a
	}
	return a - b
}

func main() {
	// result: [1,2,3,4]
	// arr := []int{1,2,3,4,5}
	// k := int(4)
	// x := int(3)

	// result: [1,2,3,4]
	// arr := []int{1,2,3,4,5}
	// k := int(4)
	// x := int(-1)

	// result: [1,2,8,9]
	// arr := []int{1,2,8,9}
	// k := int(4)
	// x := int(3)

	// result: [1,2]
	// arr := []int{1,2,8,9,10}
	// k := int(2)
	// x := int(3)

	// result: [8,9,10]
	// arr := []int{1,2,8,9,10}
	// k := int(3)
	// x := int(7)

	// result: [1]
	arr := []int{1,3}
	k := int(1)
	x := int(2)

	// result: 
	// arr := []int{}
	// k := int(0)
	// x := int(0)

	result := findClosestElements(arr, k, x)
	fmt.Printf("result = %v\n", result)
}

