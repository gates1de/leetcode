package main

import (
	"fmt"
)

func pyramidTransition(bottom string, allowed []string) bool {
	n := len(bottom)
	m := [7][7]int{}
	seen := make(map[int]bool)

	for _, a := range allowed {
		m[a[0] - 'A'][a[1] - 'A'] |= 1 << (a[2] - 'A')
	}

	alphatbets := make([][]int, n)
	for i := range n {
		alphatbets[i] = make([]int, n)
	}

	t := int(0)
	for _, char := range bottom {
		alphatbets[n - 1][t] = int(char - 'A')
		t++
	}

	return helper(m, seen, alphatbets, 0, n - 1, 0)
}

func helper(m [7][7]int, seen map[int]bool, alphabets [][]int, row int, n int, i int) bool {
	if n == 1 && i == 1 {
		return true
	} else if i == n {
		if seen[row] {
			return false
		}
		seen[row] = true
		return helper(m, seen, alphabets, 0, n - 1, 0)
	}

	w := m[alphabets[n][i]][alphabets[n][i + 1]]
	for b := range 7 {
		if ((w >> b) & 1) != 0 {
			alphabets[n - 1][i] = b
			if helper(m, seen, alphabets, row * 8 + (b + 1), n, i + 1) {
				return true
			}
		}
	}

	return false
}

func main() {
	// result: true
	// bottom := "BCD"
	// allowed := []string{"BCC","CDE","CEA","FFF"}

	// result: false
	bottom := "AAAA"
	allowed := []string{"AAB","AAC","BCD","BBE","DEF"}

	// result: false
	// bottom := ""
	// allowed := []string{}

	result := pyramidTransition(bottom, allowed)
	fmt.Printf("result = %v\n", result)
}

