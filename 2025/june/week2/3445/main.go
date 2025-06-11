package main
import (
	"fmt"
	"math"
)

func maxDifference(s string, k int) int {
	result := math.MinInt32
	sChars := []byte{'0', '1', '2', '3', '4'}

	for _, a := range sChars {
		for _, b := range sChars {
			if a == b {
				continue
			}

			best := make([]int, 4)
			for i := range best {
				best[i] = math.MaxInt32
			}

			countA := int(0)
			countB := int(0)
			prevA := int(0)
			prevB := int(0)
			left := int(-1)

			for i, char := range s {
				if byte(char) == a {
					countA++
				}
				if byte(char) == b {
					countB++
				}

				for i - left >= k && countB - prevB >= 2 {
					leftStatus := getStatus(prevA, prevB)

					if prevA - prevB < best[leftStatus] {
						best[leftStatus] = prevA - prevB
					}

					left++

					if s[left] == a {
						prevA++
					}
					if s[left] == b {
						prevB++
					}
				}

				rightStatus := getStatus(countA, countB)

				if best[rightStatus ^ 0b10] != math.MaxInt32 {
					current := (countA - countB) - best[rightStatus ^ 0b10]
					result = max(result, current)
				}
			}
		}
	}

	return result
}

func getStatus(a, b int) int {
	return ((a & 1) << 1) | (b & 1)
}

func max(a, b int) int {
	if b > a {
		return b
	}
	return a
}

func main() {
	// result: -1
	// s := "12233"
	// k := int(4)

	// result: 1
	// s := "1122211"
	// k := int(3)

	// result: -1
	s := "110"
	k := int(3)

	// result: 
	// s := ""
	// k := int(0)

	result := maxDifference(s, k)
	fmt.Printf("result = %v\n", result)
}

