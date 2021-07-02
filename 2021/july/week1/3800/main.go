package main
import (
	"fmt"
	"sort"
)

func findClosestElements(arr []int, k int, x int) []int {
	arrLen := len(arr)
	if k == len(arr) {
		return arr
	}

	if x <= arr[0] {
		return arr[:k]
	}
	if arr[arrLen - 1] <= x {
		return arr[arrLen - k:]
	}

	result := make([]int, k)
	resultIndex := int(0)
	leftIndex := int(-1)
	rightIndex := arrLen
	for i, num := range arr {
		if num == x {
			result[resultIndex] = x
			resultIndex++
		} else if num < x {
			leftIndex = max(leftIndex, i)
		} else if num > x {
			rightIndex = min(rightIndex, i)
			break
		}
	}

	left := max(0, leftIndex - k + 1)
	right := min(arrLen - 1, rightIndex + k - 1)
	for resultIndex < k {
		// fmt.Printf("result = %v, leftIndex = %v, rightIndex = %v\n", result, leftIndex, rightIndex)
		if leftIndex == rightIndex {
			result[resultIndex] = x
			resultIndex++
			leftIndex--
			rightIndex++
			continue
		}

		leftDiff := int(100000)
		rightDiff := int(100000)
		if left <= leftIndex {
			leftDiff = x - arr[leftIndex]
		}
		if rightIndex <= right {
			rightDiff = arr[rightIndex] - x
		}

		if leftDiff <= rightDiff {
			result[resultIndex] = arr[leftIndex]
			resultIndex++
			leftIndex--
		} else if leftDiff > rightDiff {
			result[resultIndex] = arr[rightIndex]
			resultIndex++
			rightIndex++
		}
	}

	sort.Slice(result, func (a, b int) bool { return result[a] < result[b] })
	return result
}

func min(a, b int) int {
	if b < a {
		return b
	}
	return a
}

func max(a, b int) int {
	if b > a {
		return b
	}
	return a
}

func main() {
	// result: [1,2,3,4]
	// arr := []int{1, 2, 3, 4, 5}
	// k := int(4)
	// x := int(3)

	// result: [1,2,3,4]
	// arr := []int{1, 2, 3, 4, 5}
	// k := int(4)
	// x := int(-1)

	// result: [4,7,8,9]
	// arr := []int{1, 2, 3, 4, 7, 8, 9}
	// k := int(4)
	// x := int(7)

	// result: [1]
	// arr := []int{1}
	// k := int(1)
	// x := int(0)

	// result: [3,3,4]
	arr := []int{0, 0, 1, 2, 3, 3, 4, 7, 7, 8}
	k := int(3)
	x := int(5)

	// result: 
	// arr := []int{}
	// k := int(0)
	// x := int(0)

	result := findClosestElements(arr, k, x)
	fmt.Printf("result = %v\n", result)
}

