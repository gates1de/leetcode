package main

import (
	"fmt"
)

func maxLength(nums []int) int {
	result := 0
	for left := range nums {
		product := int64(1)
		currentGCD := 0
		currentLCM := 1

		for right := left; right < len(nums); right++ {
			value := nums[right]
			product *= int64(value)
			if product > 25200 {
				break
			}

			currentGCD = gcd(currentGCD, value)
			currentLCM = currentLCM / gcd(currentLCM, value) * value
			if product == int64(currentGCD*currentLCM) {
				result = max(result, right-left+1)
			}
		}
	}

	return result
}

func gcd(a int, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

func main() {
	// result: 5
	// nums := []int{1,2,1,2,1,1,1}

	// result: 3
	// nums := []int{2,3,4,5,6}

	// result: 5
	nums := []int{1, 2, 3, 1, 4, 5, 1}

	// result: 0
	// nums := []int{}

	result := maxLength(nums)
	fmt.Printf("result = %v\n", result)
}
