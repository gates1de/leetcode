package main
import (
	"fmt"
	"math"
)

// https://leetcode.com/problems/find-the-shortest-superstring/discuss/397004/golang-beats-100-16ms-dp
func shortestSuperstring(words []string) string {
    var n int = len(words)
    adj_map := make([][]int, n)
    for i:=0; i<n; i++ {
        adj_map[i] = make([]int, n)
        for j:=0; j<n; j++ {
            aj, ai := []byte(words[j]), []byte(words[i])
            adj_map[i][j] = len(aj)
            for k:=1; k<=min(len(ai), len(aj)); k++ {
                if string(ai[len(ai)-k:]) == string(aj[:k]) {
                    adj_map[i][j] = len(aj) - k
                }
            }
        }
    }

    dp := make([][]int, 1<<uint(n))
    parent := make([][]int, 1<<uint(n))

    for i:=0; i<(1<<uint(n)); i++ {
        dp[i] = make([]int, n)
        parent[i] = make([]int, n)
        for j:=0; j<n; j++ {
            dp[i][j] = math.MaxInt32
            parent[i][j] = -1
        }
    }

    for i:=0; i<n; i++ {
        dp[1<<uint(i)][i] = len([]byte(words[i]))
    }

    for s:=1; s<(1<<uint(n)); s++ {
        for j:=0; j<n; j++ {
            if s & (1<<uint(j)) == 0 { continue }
            ps := s & ^(1<<uint(j))
            for i:=0; i<n; i++ {
                if dp[ps][i] + adj_map[i][j] < dp[s][j] {
                    dp[s][j] = dp[ps][i] + adj_map[i][j]
                    parent[s][j] = i;
                }
            }
        }
    }

    j, min := 0, math.MaxInt32
    for i:=0; i<n; i++ {
        if dp[(1<<uint(n))-1][i] < min {
            min = dp[(1<<uint(n))-1][i]
            j = i
        }
    }

    s := (1<<uint(n)) -1
    ans := ""
    for s!=0 {
        i := parent[s][j]
        if i < 0 {
            ans = words[j] + ans
        } else {
            adj := []byte(words[j])
            ans = string(adj[len(adj)-adj_map[i][j]:]) + ans
        }
        s &= ^(1<<uint(j))
        j = i
    }
    return ans
}


func min(a, b int) int {
    if a < b { return a }
    return b
}

// Timeout Limit Excceeded
func ngSolution(words []string) string {
	superstrings := []string{}
	for i, word := range words {
		newWords := []string{}
		newWords = append(newWords, words[:i]...)
		newWords = append(newWords, words[i + 1:]...)
		helper(newWords, word, &superstrings)
	}
	// fmt.Printf("superstrings = %v\n", superstrings)

	result := ""
	for _, superstring := range superstrings {
		if len(result) == 0 || len(result) > len(superstring) {
			result = superstring
		}
	}
	return result
}

func helper(words []string, current string, superstrings *[]string) {
	if len(words) == 0 {
		// fmt.Printf("current = %v\n", current)
		*superstrings = append(*superstrings, current)
		return
	}

	for i, word := range words {
		newWords := []string{}
		newWords = append(newWords, words[:i]...)
		newWords = append(newWords, words[i + 1:]...)
		superstring := generateSuperstring(current, word)
		// fmt.Printf("superstring = %v\n", superstring)
		helper(newWords, superstring, superstrings)
	}
}

func generateSuperstring(base string, target string) string {
	result := base + target
	for i, _ := range target {
		if len(base) - 1 < i {
			break
		}
		baseSuffix := base[len(base) - 1 - i:]
		targetPrefix := target[:i + 1]
		// fmt.Printf("baseSuffix = %v, targetPrefix = %v\n", baseSuffix, targetPrefix)
		if baseSuffix == targetPrefix {
			result = base + target[i + 1:]
		}
	}
	return result
}

func main() {
	// result: "alexlovesleetcode"
	// words := []string{"alex", "loves", "leetcode"}

	// result: "gctaagttcatgcatc"
	words := []string{"catg", "ctaagt", "gcta", "ttca", "atgcatc"}

	// result: 
	// words := []string{}

	result := shortestSuperstring(words)
	fmt.Printf("result = %v\n", result)
}

