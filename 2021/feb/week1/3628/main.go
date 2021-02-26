package main
import (
	"fmt"
)

func findLHS(nums []int) int {
	result := int(0)
	for _, num := range nums {
		sameNumCount := int(0)
		subseqCountDown := int(0)
		subseqCountUp := int(0)
		for _, n := range nums {
			// fmt.Printf("num = %v, n = %v\n", num, n)
			if num == n {
				sameNumCount++
			} else if num - n == -1 {
				subseqCountDown++
			} else if num - n == 1 {
				subseqCountUp++
			}
		}

		// fmt.Printf("same = %v, down = %v, up = %v\n", sameNumCount, subseqCountDown, subseqCountUp)
		if subseqCountDown == 0 && subseqCountUp == 0 {
			continue
		}

		subseqCount := max(subseqCountDown, subseqCountUp) + sameNumCount
		if result < subseqCount {
			result = subseqCount
		}
	}
	return result
}

func max(a int, b int) int {
    if a < b {
        return b
    }
    return a
}

func abs(a int, b int) int {
    if a - b < 0 {
        return b - a
    }
    return a - b
}

func main() {
	// nums := []int{1, 3, 2, 2, 5, 2, 3, 7}
	// nums := []int{1, 2, 3, 4}
	// nums := []int{1, 1, 1, 1}
	result := findLHS(nums)
	fmt.Printf("result = %v\n", result)
}

