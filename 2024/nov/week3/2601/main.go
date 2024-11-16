package main
import (
	"fmt"
	"math"
)

func primeSubOperation(nums []int) bool {
	maxElement := getMaxElement(nums)
	sieve := make([]bool, maxElement + 1)
	fill(sieve, true)
	sieve[1] = false

	for i := 2; i <= int(math.Sqrt(float64(maxElement + 1))); i++ {
		if !sieve[i] {
			continue
		}

		for j := i * i; j <= maxElement; j += i {
			sieve[j] = false
		}
	}

	current := int(1)
	i := int(0)
	for i < len(nums) {
		diff := nums[i] - current

		if diff < 0 {
			return false
		}

		if sieve[diff] || diff == 0 {
			i++
		}
		current++
	}

	return true
}

func getMaxElement(nums []int) int {
	result := int(-1)
	for _, num := range nums {
		if num > result {
			result = num
		}
	}
	return result
}

func fill(arr []bool, val bool) {
	for i, _ := range arr {
		arr[i] = val
	}
}

func main() {
	// result: true
	// nums := []int{4,9,6,10}

	// result: true
	// nums := []int{6,8,11,12}

	// result: false
	nums := []int{5,8,3}

	// result: false
	// nums := []int{}

	result := primeSubOperation(nums)
	fmt.Printf("result = %v\n", result)
}

