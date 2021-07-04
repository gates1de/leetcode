package main
import (
	"fmt"
)

const modulo = int(1e9 + 7)
var vowels = []string{"a", "e", "i", "o", "u"}

func countVowelPermutation(n int) int {
	var mod int64 = 1000000007
	cache := [5]int64{1, 1, 1, 1, 1}
	for i := 1; i < n; i++ {
		cacheCopy := [5]int64{}
		cacheCopy[0] += cache[1] % 1000000007
		cacheCopy[1] += (cache[0] + cache[2]) % 1000000007
		cacheCopy[2] += (cache[0] + cache[1] + cache[3] + cache[4]) % 1000000007
		cacheCopy[3] += (cache[2] + cache[4]) % 1000000007
		cacheCopy[4] += cache[0] % 1000000007
		copy(cache[:], cacheCopy[:])
	}

	var result int64
	for _, counter := range cache {
		result += counter
	}
	return int(result % mod)
}

// Slow
func mySolution(n int) int {
	if n == 1 {
		return 5
	}

	// fmt.Printf("vowels = %v\n", vowels)
	m := map[string]int{}
	for _, v := range vowels {
		m[v]++
	}

	result := int(0)
	for i := 1; i < n; i++ {
		// fmt.Printf("m = %v\n", m)
		result = count(m)
	}
	return result % modulo
}

func count(m map[string]int) int {
	copyM := map[string]int{}
	for v, numOfV := range m {
		copyM[v] = numOfV
	}

	result := int(0)
	for v, numOfV := range copyM {
		// fmt.Printf("v = %v, numOfV = %v\n", v, numOfV)
		numOfV %= modulo
		if v == "a" {
			m["e"] += numOfV
			result += numOfV
		} else if v == "e" {
			m["a"] += numOfV
			m["i"] += numOfV
			result += numOfV * 2
		} else if v == "i" {
			m["a"] += numOfV
			m["e"] += numOfV
			m["o"] += numOfV
			m["u"] += numOfV
			result += numOfV * 4
		} else if v == "o" {
			m["i"] += numOfV
			m["u"] += numOfV
			result += numOfV * 2
		} else if v == "u" {
			m["a"] += numOfV
			result += numOfV
		}
	}
	m["a"] -= copyM["a"]
	m["e"] -= copyM["e"]
	m["i"] -= copyM["i"]
	m["o"] -= copyM["o"]
	m["u"] -= copyM["u"]
	return result
}

func main() {
	// result: 5
	// n := int(1)

	// result: 10
	// n := int(2)

	// result: 68
	// n := int(5)

	// result: 759959057
	n := int(20000)

	// result: 
	// n := int(0)

	result := countVowelPermutation(n)
	fmt.Printf("result = %v\n", result)
}

