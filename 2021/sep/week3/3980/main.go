package main
import (
	"fmt"
)

func numDistinct(s string, t string) int {
	result := int(0)
	newS := ""
	m := map[rune]bool{}
	for _, tChar := range t {
		m[tChar] = true
	}
	for _, sChar := range s {
		if m[sChar] {
			newS += string(sChar)
		}
	}

	for i, sChar := range newS {
		if rune(t[0]) == sChar {
			helper(newS[i + 1:], t[1:], &result)
		}
	}
	return result
}

func helper(s string, t string, result *int) {
	// fmt.Printf("s = %v, t = %v\n", s, t)
	if len(s) == 0 || len(t) == 0 {
		return
	}

	tIndex := int(0)
	current := rune(t[tIndex])
	counts := make([]int, len(t))
	TOP:
	for _, sChar := range s {
		// fmt.Printf("current = %v, sChar = %v\n", current, sChar)
		if current != sChar && tIndex < len(t) - 1 {
			pre := current
			for current == pre {
				tIndex++
				if tIndex >= len(t) {
					break TOP
				}
				current = rune(t[tIndex])
				if current == pre {
					counts[tIndex] = counts[tIndex - 1]
				}
			}
		}

		if current == sChar {
			counts[tIndex]++
		}
	}

	// fmt.Printf("counts = %v\n", counts)
	sumCount := int(1)
	r := int(1)
	preChar := rune(' ')
	for i, count := range counts {
		currentChar := rune(t[i])
		if preChar == currentChar {
			r++
			continue
		} else {
			r = 1
			preChar = currentChar
		}
		sumCount *= combination(count, r)
	}
	*result += sumCount
}

func combination(n int, r int) int {
	result := int(1)
	for i := n; i > n - r; i-- {
		result *= i
	}
	for i := 1; i <= r; i++ {
		result /= i
	}
	return result
}

func main() {
	// result: 3
	// s := "rabbbit"
	// t := "rabbit"

	// result: 5
	s := "babgbag"
	t := "bag"

	// result: 
	// s := ""
	// t := ""

	// result: 
	// s := ""
	// t := ""

	result := numDistinct(s, t)
	fmt.Printf("result = %v\n", result)
}

