package main
import (
	"fmt"
	"strings"
)

// Past solution (2021/jan/week4/3619)
func getSmallestString(n int, k int) string {
	result := ""
	for k > 0 {
		if n == k {
			result = strings.Repeat("a", n) + result
			break
		}

		if k - n >= 26 {
			zCount := (k - n) / 26
			result += strings.Repeat("z", zCount)
			k -= 26 * zCount
			n -= zCount
			if n >= k {
				k += 26
				n++
			}
			continue
		}

		c := string(rune(97 + k - n))
		result = c + result
		n--
		k = n
	}
	return result
}

// Slow
func mySolution(n int, k int) string {
	return helper("", n, k)
}

func helper(current string, n int, k int) string {
	// fmt.Println(current, n, k)
	if n == 0 {
		return current
	}

	if n == k {
		current = "a" + current
		return helper(current, n - 1, k - 1)
	}

	if k < n + 26 {
		alphaNum := k - (n - 1)
		current = string(rune(byte(96 + alphaNum))) + current
		return helper(current, n - 1, k - alphaNum)
	}

	current = current + "z"
	return helper(current, n - 1, k - 26)
}

func main() {
	// result: "aay"
	// n := int(3)
	// k := int(27)

	// result: "aaszz"
	// n := int(5)
	// k := int(73)

	// result: "a"
	// n := int(1)
	// k := int(1)

	// result: "aayzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzz"
	// n := int(85)
	// k := int(2159)

	// result: omit
	n := int(91713)
	k := int(116185)

	// result: 
	// n := int(0)
	// k := int(0)

	result := getSmallestString(n, k)
	fmt.Printf("result = %v\n", result)
}

