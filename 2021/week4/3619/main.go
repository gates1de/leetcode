package main
import (
	"fmt"
)

func getSmallestString(n int, k int) string {
	result := ""
	for k > 0 {
		if n == k {
			for i := 0; i < n; i++ {
				result = "a" + result
			}
			break
		}

		if k - n < 26 {
			c := string(rune(97 + k - n))
			// fmt.Printf("c = %v\n", c)
			result = c + result
			n--
			k = n
			continue
		}
		result += "z"
		k -= 26
		n--
	}
	return result
}

func main() {
	// n := int(3)
	// k := int(27)

	// n := int(5)
	// k := int(73)

	n := int(91713)
	k := int(116185)

	result := getSmallestString(n, k)
	fmt.Printf("result = %v\n", result)
}

