package main
import (
	"fmt"
	"math"
)

func grayCode(n int) []int {
	numOfCodes := int(math.Pow(2.0, float64(n)))
	result := make([]int, numOfCodes)
	// fmt.Printf("numOfCodes = %v\n", numOfCodes)
	m := map[int]bool{}
	m[0] = true
	for i, _ := range result {
		if i == 0 {
			continue
		}
		current := result[i - 1]
		// bin := strconv.FormatInt(int64(target), 2)
		// fmt.Printf("target = %v, bin = %v\n", target, bin)

		for j := 0; j <= n; j++ {
			twoPow := int(math.Pow(2.0, float64(j)))
			target := current + twoPow
			minusTarget := current - twoPow

			// fmt.Printf("isOneBitDiff = %v, minus = %v\n", isOneBitDiff(current, target), isOneBitDiff(current, minusTarget))
			if isOneBitDiff(current, target) && target <= numOfCodes && !m[target] {
				result[i] = target
				m[target] = true
				break
			}
			if isOneBitDiff(current, minusTarget) && 0 <= minusTarget && !m[minusTarget] {
				result[i] = minusTarget
				m[minusTarget] = true
				break
			}
		}
	}
	return result
}

func isOneBitDiff(a, b int) bool {
	// REF: https://www.geeksforgeeks.org/check-whether-two-numbers-differ-one-bit-position/
	x := a ^ b
	return x != 0 && ((x & (x - 1)) == 0)
}

func main() {
	// result: [0,1,3,2]
	// n := int(2)

	// result: [0,1]
	// n := int(1)

	// result: [0,1,3,2,6,7,5,4,12,13,15,14,10,11,9,8]
	// n := int(4)

	// result: abbreviation
	n := int(16)

	// result: 
	// n := int(0)

	result := grayCode(n)
	fmt.Printf("result = %v\n", result)
}

