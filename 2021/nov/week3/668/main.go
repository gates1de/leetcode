package main
import (
	"fmt"
	"sort"
)

func findKthNumber(m int, n int, k int) int {
	if m * n == k {
		return m * n
	}

	if m < n {
		m, n = n, m
	}

	counter := map[int]int{}
	for i := 1; i <= m; i++ {
		for j := 1; j <= i; j++ {
			if j <= n {
				counter[i * j]++
			}
			if i != j && i <= n {
				counter[i * j]++
			}
		}
	}

	sum := int(0)
	keys := []int{}
	for key, _ := range counter {
		keys = append(keys, key)
	}

	sort.Ints(keys)
	for _, key := range keys {
		sum += counter[key]
		if sum >= k {
			return key
		}
	}
	return 0
}

func main() {
	// result: 3
	// m := int(3)
	// n := int(3)
	// k := int(5)

	// result: 6
	// m := int(2)
	// n := int(3)
	// k := int(6)

	// result: 2
	// m := int(1)
	// n := int(3)
	// k := int(2)

	// result: 211005
	m := int(3000)
	n := int(3000)
	k := int(1000000)

	// result: 
	// m := int(0)
	// n := int(0)
	// k := int(0)

	result := findKthNumber(m, n, k)
	fmt.Printf("result = %v\n", result)
}

