package main
import (
	"fmt"
)

func zeroFilledSubarray(nums []int) int64 {
	zeroCountMap := make(map[int]int)
	maxZeroCount := int(0)
	zeroCount := int(0)
	for _, num := range nums {
		if num == 0 {
			zeroCount++
		} else {
			zeroCountMap[zeroCount]++
			zeroCount = 0
		}

		if zeroCount > maxZeroCount {
			maxZeroCount = zeroCount
		}
	}
	zeroCountMap[zeroCount]++

	result := int64(0)
	for i := 1; i <= maxZeroCount; i++ {
		count := zeroCountMap[i]
		if count == 0 {
			continue
		}

		addVal := int(0)
		for j := 1; j <= i; j++ {
			addVal += j
		}
		result += int64(count * addVal)
	}

	return result
}

func main() {
	// result: 6
	// nums := []int{1,3,0,0,2,0,0,4}

	// result: 9
	// nums := []int{0,0,0,2,0,0}

	// result: 0
	nums := []int{2,10,2019}

	// result: 
	// nums := []int{}

	result := zeroFilledSubarray(nums)
	fmt.Printf("result = %v\n", result)
}

