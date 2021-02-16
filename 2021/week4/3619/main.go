package main
import (
	"fmt"
	"strings"
)

func getSmallestString(n int, k int) string {
	result := ""
	for k > 0 {
		if n == k {
			result = strings.Repeat("a", n) + result
			break
		}

		if k - n >= 26 {
			zCount := (k - n) / 26
			// fmt.Printf("zCount = %v\n", zCount)
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
		// fmt.Printf("c = %v, 97 + k - n = %v\n", c, 97 + k - n)
		result = c + result
		n--
		k = n
	}
	return result
}

func main() {
	// n := int(3)
	// k := int(27)

	// n := int(5)
	// k := int(73)

	// n := int(85)
	// k := int(2159)

	n := int(91713)
	k := int(116185)

	result := getSmallestString(n, k)
	fmt.Printf("result = %v\n", result)
}

