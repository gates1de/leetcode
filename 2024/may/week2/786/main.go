package main
import (
	"fmt"
)

func kthSmallestPrimeFraction(arr []int, k int) []int {
	n := len(arr)
	left := float32(0)
	right := float32(1.0)

	for left < right {
		mid := (left + right) / float32(2)

		maxFraction := float32(0)
		totalSmallerFractions := int(0)
		numeratorIndex := int(0)
		denominatorIndex := int(0)
		j := int(1)

		for i := 0; i < n - 1; i++ {
			for j < n && (float32(arr[i]) >= mid * float32(arr[j])) {
				j++
			}

			totalSmallerFractions += n - j

			if j == n {
				break
			}

			fraction := float32(arr[i]) / float32(arr[j])

			if fraction > maxFraction {
				numeratorIndex = i
				denominatorIndex = j
				maxFraction = fraction
			}
		}

		if totalSmallerFractions == k {
			return []int{arr[numeratorIndex], arr[denominatorIndex]}
		} else if (totalSmallerFractions > k) {
			right = mid
		} else {
			left = mid
		}
	}

	return []int{}
}

func main() {
	// result: [2,5]
	// arr := []int{1,2,3,5}
	// k := int(3)

	// result: [1,7]
	arr := []int{1,7}
	k := int(1)

	// result: []
	// arr := []int{}
	// k := int(0)

	result := kthSmallestPrimeFraction(arr, k)
	fmt.Printf("result = %v\n", result)
}

