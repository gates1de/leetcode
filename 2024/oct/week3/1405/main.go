package main
import (
	"fmt"
)

func longestDiverseString(a int, b int, c int) string {
	aSeq := int(0)
	bSeq := int(0)
	cSeq := int(0)
	result := ""
	n := a + b + c

	for i := 0; i < n; i++ {
		if (a >= b && a >= c && aSeq != 2) || (a > 0 && (bSeq == 2 || cSeq == 2)) {
			result += "a"
			a--
			aSeq++
			bSeq = 0
			cSeq = 0
		} else if (b >= a && b >= c && bSeq != 2) || (b > 0 && (aSeq == 2 || cSeq == 2)) {
			result += "b"
			b--
			bSeq++
			aSeq = 0
			cSeq = 0
		} else if (c >= a && c >= b && cSeq != 2) || (c > 0 && (aSeq == 2 || bSeq == 2)) {
			result += "c"
			c--
			cSeq++
			aSeq = 0
			bSeq = 0
		}
	}

	return result
}

func main() {
	// result: "ccaccbcc"
	// a := int(1)
	// b := int(1)
	// c := int(7)

	// result: "aabaa"
	a := int(7)
	b := int(1)
	c := int(0)

	// result: ""
	// a := int(0)
	// b := int(0)
	// c := int(0)

	result := longestDiverseString(a, b, c)
	fmt.Printf("result = %v\n", result)
}

