package main
import (
	"fmt"
	"sort"
	"container/heap"

	localHeap "heap/heap"
)

func getLengthOfOptimalCompression(s string, k int) int {
	n := len(s)
	dp := make([][]int, 101)
	for i, _ := range dp {
		dp[i] = make([]int, 101)
		for j, _ := range dp[i] {
			dp[i][j] = 10000
		}
	}

	dp[0][0] = 0
	for i := 1; i <= n; i++ {
		for j := 0; j <= k; j++ {
			count := 0
			del := 0
			for l := i; l >= 1; l-- {
				if s[l - 1] == s[i - 1] {
					count++
				} else {
					del++
				}

				if j - del >= 0 {
					numsLength := int(0)
					if count >= 100 {
						numsLength = 3
					} else if count >= 10 {
						numsLength = 2
					} else if count >= 2 {
						numsLength = 1
					}
					dp[i][j] = min(dp[i][j], dp[l - 1][j - del] + 1 + numsLength)
				}
			}

			if j > 0 {
				dp[i][j] = min(dp[i][j], dp[i - 1][j - 1])
			}
		}
	}

	return dp[n][k]
}

func min(a, b int) int {
	if b < a {
		return b
	}
	return a
}

// Wrong Answer
func ngSolution(s string, k int) int {
	if len(s) == 1 {
		if k >= 1 {
			return 0
		}
		return 1
	} else if len(s) <= k {
		return 0
	}

	chars := []byte(s)
	sort.Slice(chars, func(a, b int) bool {
		return chars[a] < chars[b]
	})
	s = string(chars)
	s += " "

	compHeap := &localHeap.CompHeap{}
	heap.Init(compHeap)
	pre := s[0]
	count := int(1)
	for _, char := range s[1:] {
		b := byte(char)
		if pre == b {
			count++
			continue
		}

		heap.Push(compHeap, localHeap.Comp{Char: byte(pre), Count: count})
		pre = b
		count = 1
	}

	compressedS := ""
	for compHeap.Len() > 0 {
		comp := heap.Pop(compHeap).(localHeap.Comp)
		fmt.Println(comp)

		if k == 0 {
			if comp.Count == 1 {
				compressedS += string(comp.Char)
			} else {
				compressedS += fmt.Sprintf("%v%v", string(comp.Char), comp.Count)
			}
			continue
		}

		if k < comp.Count {
			if compHeap.Len() > 0 {
				nextComp := heap.Pop(compHeap).(localHeap.Comp)
				if comp.Count < 10 && nextComp.Count >= 10 && nextComp.Count - k < 10 {
					if nextComp.Count - k == 1 {
						compressedS += string(nextComp.Char)
					} else {
						compressedS += fmt.Sprintf("%v%v", string(nextComp.Char), nextComp.Count - k)
					}

					if comp.Count == 1 {
						compressedS += string(comp.Char)
					} else {
						compressedS += fmt.Sprintf("%v%v", string(comp.Char), comp.Count)
					}
				} else {
					if comp.Count - k == 1 {
						compressedS += string(comp.Char)
					} else {
						compressedS += fmt.Sprintf("%v%v", string(comp.Char), comp.Count - k)
					}

					if nextComp.Count == 1 {
						compressedS += string(nextComp.Char)
					} else {
						compressedS += fmt.Sprintf("%v%v", string(nextComp.Char), nextComp.Count)
					}
				}

				k = 0
			} else {
				compressedS += fmt.Sprintf("%v%v", string(comp.Char), comp.Count - k)
				k = 0
			}
		} else {
			k -= comp.Count
		}
	}

	fmt.Println(compressedS)
	return len(compressedS)
}

func main() {
	// result: 4
	// s := "aaabcccd"
	// k := int(2)

	// result: 2
	// s := "aabbaa"
	// k := int(2)

	// result: 3
	// s := "aaaaaaaaaaa"
	// k := int(0)

	// result: 4
	// s := "aaaaaaaaaaabcccd"
	// k := int(4)

	// result: 1
	s := "abc"
	k := int(2)

	// result: 
	// s := ""
	// k := int(0)

	result := getLengthOfOptimalCompression(s, k)
	fmt.Printf("result = %v\n", result)
}

