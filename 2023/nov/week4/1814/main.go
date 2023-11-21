package main
import (
	"fmt"
)

const modulo = int(1e9 + 7)

func countNicePairs(nums []int) int {
	arr := make([]int, len(nums))
	for i, num := range nums {
		arr[i] = num - reverse(num)
	}

	m := make(map[int]int)
	result := int(0)
	for _, num := range arr {
		result = (result + m[num]) % modulo
		m[num] += 1
	}

	return result
}

func reverse(num int) int {
	result := int(0)
	for num > 0 {
		result = result * 10 + num % 10
		num /= 10
	}
	return result
}

func main() {
	// result: 2
	// nums := []int{42,11,1,97}

	// result: 4
	nums := []int{13,10,35,24,76}

	// result: 
	// nums := []int{}

	result := countNicePairs(nums)
	fmt.Printf("result = %v\n", result)
}

