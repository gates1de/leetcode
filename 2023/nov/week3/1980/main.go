package main
import (
	"fmt"
	"sort"
	"strconv"
)

func findDifferentBinaryString(nums []string) string {
	n := len(nums)
	maxNum := int(1)
	for i := 0; i < n; i++ {
		maxNum *= 2
	}

	sort.Slice(nums, func(a, b int) bool {
		numA, _ := strconv.ParseInt(nums[a], 2, 32)
		numB, _ := strconv.ParseInt(nums[b], 2, 32)
		return numA < numB
	})

	for i, bin := range nums {
		num, _ := strconv.ParseInt(bin, 2, 32)
		if i != int(num) {
			return fmt.Sprintf("%016b", i)[16 - n:]
		}
	}

	return fmt.Sprintf("%016b", maxNum - 1)[16 - n:]
}

func main() {
	// result: "11"
	// nums := []string{"01","10"}

	// result: "11"
	// nums := []string{"00","01"}

	// result: "101"
	nums := []string{"111","011","001"}

	// result: ""
	// nums := []string{}

	result := findDifferentBinaryString(nums)
	fmt.Printf("result = %v\n", result)
}

