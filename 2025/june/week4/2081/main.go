package main
import (
	"fmt"
)

func kMirror(k int, n int) int64 {
	digit := make([]int, 100)
	left := int(1)
	count := int(0)
	result := int64(0)

	for count < n {
		right := left * 10

		for op := 0; op < 2; op++ {
			for i := left; i < right && count < n; i++ {
				combined := int64(i)
				x := i

				if op == 0 {
					x = i / 10
				}

				for x > 0 {
					combined = combined * 10 + int64(x % 10)
					x /= 10
				}

				if isPalindrome(combined, k, digit) {
					count++
					result += combined
				}
			}
		}

		left = right
	}

	return result
}

func isPalindrome(x int64, k int, digit []int) bool {
	length := int(-1)

	for x > 0 {
		length++
		digit[length] = int(x % int64(k))
		x /= int64(k)
	}

	for i, j := 0, length; i < j; i, j = i + 1, j - 1 {
		if digit[i] != digit[j] {
			return false
		}
	}

	return true
}

func main() {
	// result: 25
	// k := int(2)
	// n := int(5)

	// result: 499
	// k := int(3)
	// n := int(7)

	// result: 20379000
	k := int(7)
	n := int(17)

	// result: 
	// k := int(0)
	// n := int(0)

	result := kMirror(k, n)
	fmt.Printf("result = %v\n", result)
}

