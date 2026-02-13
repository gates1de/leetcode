package main

import (
	"fmt"
	"math"
)

func longestBalanced(s string) int {
	return max(longestSingle(s), longestAbc(s))
}

func longestSingle(s string) int {
	a, b, c, ma := 0, 0, 0, 0
	for i := range len(s) {
		switch s[i] {
		case 'a':
			a, b, c, ma = a + 1, 0, 0, max(ma, b, c)

		case 'b':
			a, b, c, ma = 0, b + 1, 0, max(ma, a, c)

		default:
			a, b, c, ma = 0, 0, c + 1, max(ma, a, b)
		}
	}

	return max(ma, a, b, c)
}

type Key = [2]int

func recalc(result, i int, m map[Key]int, k Key) int {
	j, exists := m[k]
	if exists {
		return max(result, i - j)
	}
	m[k] = i

	return result
}

func longestAbc(s string) int {
	m := [4]map[Key]int{
		{Key{}: -1},
		{Key{}: -1},
		{Key{}: -1},
		{Key{}: -1},
	}

	abc := [3]int{}
	result := int(0)
	for i := range len(s) {
		abc[s[i] - 'a']++
		result = recalc(result, i, m[0], Key{abc[0] - abc[1], abc[0] - abc[2]})
		result = recalc(result, i, m[1], Key{abc[0] - abc[1], abc[2]})
		result = recalc(result, i, m[2], Key{abc[1] - abc[2], abc[0]})
		result = recalc(result, i, m[3], Key{abc[2] - abc[0], abc[1]})
	}

	return result
}

func max(nums ...int) int {
	result := math.MinInt32
	for _, num := range nums {
		if result < num {
			result = num
		}
	}

	return result
}

func main() {
	// result: 4
	// s := "abbac"

	// result: 3
	// s := "zzabccy"

	// result: 2
	s := "aba"

	// result: 
	// s := ""

	result := longestBalanced(s)
	fmt.Printf("result = %v\n", result)
}

