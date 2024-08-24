package main
import (
	"fmt"
	"math"
	"strconv"
)

func nearestPalindromic(n string) string {
	num := parseInt(n)
	a := nextPalindrome(num)
	b := minBinarySearch(num)

	if abs(a - num) <= abs(b - num) {
		return fmt.Sprintf("%d", a)
	}

	return fmt.Sprintf("%d", b)
}

func convert(num int64) int64 {
	s := fmt.Sprintf("%d", num)
	chars := []byte(s)
	n := len(s)
	l := (n - 1) / 2
	r := n / 2

	for l >= 0 {
		chars[r] = chars[l]
		r++
		l--
	}

	return parseInt(string(chars))
}

func nextPalindrome(num int64) int64 {
	left := int64(0)
	right := num
	result := int64(math.MinInt64)

	for left <= right {
		mid := (right - left) / 2 + left
		palindrome := convert(mid)
		if palindrome < num {
			result = palindrome
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	return result
}

func minBinarySearch(num int64) int64 {
	left := num
	right := int64(1e18)
	result := int64(math.MinInt64)

	for left <= right {
		mid := (right - left) / 2 + left
		palindrome := convert(mid)
		if palindrome > num {
			result = palindrome
			right = mid - 1
		} else {
			left = mid + 1
		}
	}

	return result
}

func abs(num int64) int64 {
	if num < 0 {
		return -num
	}
	return num
}

func parseInt(num string) int64 {
	result, _ := strconv.ParseUint(num, 10, 64)
	return int64(result)
}

func main() {
	// result: "121"
	// n := "123"

	// result: "0"
	n := "1"

	// result: ""
	// n := ""

	result := nearestPalindromic(n)
	fmt.Printf("result = %v\n", result)
}

