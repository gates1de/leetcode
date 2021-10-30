package main
import (
	"fmt"
)

func longestDupSubstring(s string) string {
	result := ""
	n := len(s)
	l := int(0)
	r := n - 1
    for l <= r {
        m := l + (r - l) / 2
        dup := ""
        tmp := make(map[string]bool, n - m)
        for i := 0; i + m <= n; i++ {
			// fmt.Printf("%v: tmp = %v\n", i, tmp)
            if _, ok := tmp[s[i:i + m]]; !ok {
                tmp[s[i:i + m]] = true
            } else {
                dup = s[i:i + m]
                break
            }
        }

        if dup == "" {
            r = m - 1
        } else {
            l = m + 1
            result = dup
        }
    }

	if result == "" {
		m := map[rune]bool{}
		for _, c := range s {
			if m[c] {
				return string(c)
			}
			m[c] = true
		}
	}

    return result
}

// Wrong Answer
func ngSolution(s string) string {
	m := map[rune]int{}
	for _, c := range s {
		m[c]++
	}

	result := ""
	current := ""
	for i, c := range s {
		if m[c] == 1 {
			current = ""
			continue
		}

		current += string(c)
		targetStr := s[i - len(current) + 2:]
		// fmt.Printf("%v: current = %v, targetStr = %v\n", i, current, targetStr)
		for j := 0; j < len(targetStr) - len(current) + 1; j++ {
			if m[rune(targetStr[j])] == 1 {
				break
			}
			sub := targetStr[j:j + len(current)]
			// fmt.Printf("%v: sub = %v\n", j, sub)
			if current == sub && len(result) < len(current) {
				result = current
			}
		}
	}
	return result
}

func main() {
	// result: "ana"
	// s := "banana"

	// result: ""
	// s := "abcd"

	// result: "fabfa"
	// s := "wiwinfabfabfa"

	// result: "a"
	// s := "aa"

	// result: "ma"
	// s := "nnpxouomcofdjuujloanjimymadkuepightrfodmauhrsy"

	// result: "t"
	s := "zxcvdqkfawuytt"

	// result: 
	// s := ""

	result := longestDupSubstring(s)
	fmt.Printf("result = %v\n", result)
}

