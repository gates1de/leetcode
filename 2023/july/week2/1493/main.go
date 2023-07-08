package main
import (
	"fmt"
)

func longestSubarray(nums []int) int {
	n := len(nums)
	if n <= 1 {
		return 0
	}

	counter := make([]int, 1, len(nums))
	for _, num := range nums {
		lastIndex := len(counter) - 1
		if num == 1 {
			if counter[lastIndex] <= -1 {
				counter = append(counter, 1)
				continue
			}
			counter[lastIndex]++
		} else {
			if counter[lastIndex] <= -1 {
				counter[lastIndex]--
				continue
			}
			counter = append(counter, -1)
		}
	}

	if len(counter) == 1 {
		return counter[0] - 1
	} else if len(counter) == 2 {
		return max(counter[0] + counter[1], max(counter[0], counter[1]))
	}

	result := int(0)
	for i, num := range counter[:len(counter) - 2] {
		next := counter[i + 1]
		nextNext := counter[i + 2]
		if next <= -2 {
			result = max(result, max(num, nextNext))
		} else {
			result = max(result, max(num + nextNext, max(num, max(next, nextNext))))
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
	// nums := []int{1,1,0,1}

	// result: 5
	// nums := []int{0,1,1,1,0,1,1,0,1}

	// result: 2
	// nums := []int{1,1,1}

	// result: 3
	// nums := []int{0,1,1,1,0,0,1,1,0,1}

	// result: 3
	nums := []int{1,1,1,0}

	// result: 
	// nums := []int{}

	result := longestSubarray(nums)
	fmt.Printf("result = %v\n", result)
}

