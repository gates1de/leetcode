package main
import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func countGoodIntegers(n int, k int) int64 {
	m := make(map[string]bool)
	base := pow(10, (n - 1) / 2)
	skip := n & 1

	for num := base; num < base * 10; num++ {
		s := strconv.Itoa(num)
		reverseS := reverse(s)
		s += reverseS[skip:]

		palindromicInt, _ := strconv.ParseInt(s, 10, 64)
		if palindromicInt % int64(k) == 0 {
			chars := strings.Split(s, "")
			sort.Strings(chars)
			m[strings.Join(chars, "")] = true
		}
	}

	factorial := make([]int64, n + 1)
	factorial[0] = 1
	for i := 1; i <= n; i++ {
		factorial[i] = factorial[i - 1] * int64(i)
	}

	result := int64(0)
	for str, _ := range m {
		count := make([]int, 10)
		for _, char := range str {
			count[char - '0']++
		}

		total := int64(n - count[0]) * factorial[n - 1]
		for _, c := range count {
			total /= factorial[c]
		}
		result += total
	}

	return result
}

func pow(a, b int) int {
	result := int(1)
	for i := 0; i < b; i++ {
		result *= a
	}
	return result
}

func reverse(s string) string {
	chars := []rune(s)
	for i, j := 0, len(chars) - 1; i < j; i, j = i + 1, j - 1 {
		chars[i], chars[j] = chars[j], chars[i]
	}
	return string(chars)
}

func main() {
	// result: 27
	// n := int(3)
	// k := int(5)

	// result: 2
	// n := int(1)
	// k := int(4)

	// result: 2468
	n := int(5)
	k := int(6)

	// result: 
	// n := int(0)
	// k := int(0)

	result := countGoodIntegers(n, k)
	fmt.Printf("result = %v\n", result)
}

