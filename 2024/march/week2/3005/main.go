package main
import (
	"fmt"
)

func maxFrequencyElements(nums []int) int {
	m := make(map[int]int, len(nums))
	maxFreq := int(0)
	maxFreqCount := int(1)
	for _, num := range nums {
		m[num]++
		if m[num] > maxFreq {
			maxFreq = m[num]
			maxFreqCount = 1
		} else if m[num] == maxFreq {
			maxFreqCount++
		}
	}
	return maxFreq * maxFreqCount
}

func main() {
	// results: 4
	// nums := []int{1,2,2,3,1,4}

	// results: 5
	nums := []int{1,2,3,4,5}

	// results: 
	// nums := []int{}

	result := maxFrequencyElements(nums)
	fmt.Printf("result = %v\n", result)
}

