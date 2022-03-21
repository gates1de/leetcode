package main
import (
	"fmt"
)

func partitionLabels(s string) []int {
	m := map[rune][]int{}
	for i, r := range s {
		if m[r] == nil {
			m[r] = []int{-1, -1}
		}
		if m[r][0] == -1 {
			m[r][0] = i
		} else {
			m[r][1] = i
		}
	}

	// fmt.Println(m)

	result := []int{}
	left := int(-1)
	right := int(-1)
	for i := 0; i < len(s); i++ {
		r := rune(s[i])
		// fmt.Println(string(r), left, right)

		if right == -1 && m[r][1] == -1 {
			result = append(result, 1)
			left = -1
			right = -1
			continue
		}

		if left == -1 {
			left = i
		}

		if right == -1 {
			right = m[r][1]
		} else if right == i {
			result = append(result, right - left + 1)
			left = -1
			right = -1
		} else {
			if right < m[r][1] {
				right = m[r][1]
			}
		}
	}

	return result
}

func main() {
	// result: [9,7,8]
	// s := "ababcbacadefegdehijhklij"

	// result: [10]
	// s := "eccbbbbdec"

	// result: [19]
	// s := "abbcdefegdhijhklija"

	// result: [1]
	// s := "a"

	// result: [5,1,7,5]
	s := "abcbadefghijeklmnk"

	// result: 
	// s := ""

	result := partitionLabels(s)
	fmt.Printf("result = %v\n", result)
}

