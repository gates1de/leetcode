package main
import (
	"fmt"
)

func countSubstrings(s string) int {
    if len(s) <= 0 {
        return len(s)
    }

    pre := rune(s[0])
    unduplicatedS := []rune{pre}
    dup := []rune{pre}
    dups := [][]rune{}
    tmp := s + "-"
    for _, c := range tmp[1:] {
        if pre == c {
            dup = append(dup, c)
            continue
        }

        if len(dup) > 1 {
            dups = append(dups, dup)
        }

        if c != rune('-') {
            unduplicatedS = append(unduplicatedS, c)
            if len(dup) % 2 == 0 {
                unduplicatedS = append(unduplicatedS, c)
            }
        }
        dup = []rune{c}

        pre = c
    }

    result := len(s)
    for _, dup := range dups {
        for i := 1; i < len(dup); i++ {
            result += i
        }
    }

    if len(unduplicatedS) <= 2 {
        return result
    }

    for i := 1; i < len(s) - 1; i++ {
        j := i - 1
        k := i + 1
        isCounted := false
        for j >= 0 && k < len(s) {
            left := s[j]
            right := s[k]
            if left == right {
                if isCounted {
                    result++
                } else {
                    if left != s[i] {
                        isCounted = true
                        result++
                    }
                }
            } else {
                break
            }

            j--
            k++
        }

        j = i
        k = i + 1
        isCounted = false
        for j >= 0 && k < len(s) {
            left := s[j]
            right := s[k]
            if left == right {
                if isCounted {
                    result++
                } else {
                    if left != s[i] {
                        isCounted = true
                        result++
                    }
                }
            } else {
                break
            }

            j--
            k++
        }
    }

    return result
}

func main() {
	// result: 3
	// s := "abc"

	// result: 6
	s := "aaa"

	// result: 
	// s := ""

	result := countSubstrings(s)
	fmt.Printf("result = %v\n", result)
}

