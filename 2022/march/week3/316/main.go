package main
import (
	"fmt"
	"strings"
)

func removeDuplicateLetters(s string) string {
	m := map[byte]int{}
	n := len(s)
    for i := 0; i < n; i++ {
        m[s[i]]++
    }

    set := map[byte]bool{}

    var stack []byte
    for i := 0; i < n; i++ {
        m[s[i]]--
		stackLen := len(stack)
        if !set[s[i]] {
            for stackLen > 0 && stack[stackLen - 1] > s[i] && m[stack[stackLen - 1]] > 0 {
                set[stack[stackLen - 1]] = false
                stack = stack[:stackLen - 1]
				stackLen = len(stack)
            }
            stack = append(stack, s[i])
            set[s[i]] = true
        }
    }
    return string(stack)
}

// Wrong Answer
func ngSolution(s string) string {
	if len(s) <= 1 {
		return s
	}

	result := ""
	m := map[rune][]int{}
	for i, r := range s {
		result += " "
		if m[r] == nil {
			m[r] = []int{}
		}
		m[r] = append(m[r], i)
	}
	// fmt.Println(m)

	added := map[rune]int{}
	for b := byte(97); b <= 122; b++ {
		added[rune(b)] = -1
	}

	for r, indices := range m {
		if len(indices) == 1 {
			added[r] = indices[0]
			result = replace([]rune(result), r, indices[0])
			delete(m, r)
		}
	}

	// fmt.Println(added)
	for b := byte(97); b <= 122; b++ {
		r := rune(b)
		indices := m[r]
		if indices == nil {
			continue
		}

		min := len(s)
		max := int(-1)
		for b := byte(97); b <= 122; b++ {
			addedR := rune(b)
			addedIndex := added[addedR]
			if addedIndex < 0 {
				continue
			}

			if r > addedR && addedIndex < min {
				min = addedIndex
			}
			if r < addedR && addedIndex > max {
				max = addedIndex
			}
		}

		targetIndex := indices[0]
		if indices[len(indices) - 1] < min {
			targetIndex = indices[0]
		} else if max < indices[0] {
			targetIndex = indices[len(indices) - 1]
		} else {
			for _, index := range indices {
				if (min < index && index < max) || max < index {
					targetIndex = index
					break
				}
			}
		}

		result = replace([]rune(result), r, targetIndex)
		added[r] = targetIndex
	}

	return strings.ReplaceAll(result, " ", "")
}

func replace(s []rune, r rune, i int) string {
	s[i] = r
	return string(s)
}

func main() {
	// result: "abc"
	// s := "bcabc"

	// result: "acdb"
	// s := "cbacdcbc"

	// result: "a"
	// s := "a"

	// result: "x"
	// s := "xxx"

	// result: "abcdefghijklmnopqrstuvwxyz"
	s := "abcdefghijklmnopqrstuvwxyzyxwvutsrqponmlkjihgfedcba"

	// result: 
	// s := ""

	result := removeDuplicateLetters(s)
	fmt.Printf("result = %v\n", result)
}

