package main
import (
	"fmt"
)

func getHappyString(n int, k int) string {
	total := int(3 * (1 << (n - 1)))
	if k > total {
		return ""
	}

	result := make([]byte, n)

	for i, _ := range result {
		result[i] = 'a'
	}

	nextSmallest := make(map[byte]byte)
	nextSmallest['a'] = 'b'
	nextSmallest['b'] = 'a'
	nextSmallest['c'] = 'a'

	nextGreatest := make(map[byte]byte)
	nextGreatest['a'] = 'c'
	nextGreatest['b'] = 'c'
	nextGreatest['c'] = 'b'

	startA := int(1)
	startB := startA + (1 << (n - 1))
	startC := startB + (1 << (n - 1))

	if k < startB {
		result[0] = 'a'
		k -= startA
	} else if k < startC {
		result[0] = 'b'
		k -= startB
	} else {
		result[0] = 'c'
		k -= startC
	}

	for i := 1; i < n; i++ {
		mid := int(1 << (n - i - 1))

		if k < mid {
			result[i] = nextSmallest[result[i - 1]]
		} else {
			result[i] = nextGreatest[result[i - 1]]
			k -= mid
		}
	}

	return string(result)
}

func main() {
	// result: "c"
	// n := int(1)
	// k := int(3)

	// result: ""
	// n := int(1)
	// k := int(4)

	// result: "cab"
	n := int(3)
	k := int(9)

	// result: ""
	// n := int(0)
	// k := int(0)

	result := getHappyString(n, k)
	fmt.Printf("result = %v\n", result)
}

