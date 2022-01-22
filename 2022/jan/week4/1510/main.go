package main
import (
	"fmt"
	"sort"
)

const MAX = int(100000)

func winnerSquareGame(n int) bool {
    dp := make([]bool, n + 1)
    for i := 1; i <= n; i++ {
        for k := 1; k * k <= i; k++ {
            if !dp[i - k * k] {
                dp[i] = true
                break
            }
        }
    }
    return dp[n]
}

// Wrong Answer
func ngSolution(n int) bool {
	removes := []int{}
	for i := 1; i * i <= MAX; i++ {
		removes = append(removes, i * i)
	}
	removes = removes[:sort.SearchInts(removes, n + 1)]
	return helper(true, n, removes, map[[2]int]bool{})
}

func helper(isAliceTurn bool, n int, removes []int, visited map[[2]int]bool) bool {
    if n < 0 {
        return false
    }

    if isAliceTurn {
        visited[[2]int{0, n}] = true
    } else {
        visited[[2]int{-1, n}] = true
    }

    for i := len(removes) - 1; i >= 0; i-- {
        remove := removes[i]
        nextN := n - remove

        if nextN == 0 {
            return isAliceTurn
        } else if nextN < 0 {
            continue
        }

        key := [2]int{0, nextN}
        if !isAliceTurn {
            key = [2]int{-1, nextN}
        }
        if visited[key] {
            continue
        }

        newRemoves := removes[:sort.SearchInts(removes, nextN + 1)]
        isWinAlice := helper(!isAliceTurn, nextN, newRemoves, visited)

        if isAliceTurn && isWinAlice {
            return true
        }
        if !isAliceTurn && !isWinAlice {
            return false
        }
    }

    return !isAliceTurn
}

func main() {
	// result: true
	// n := int(1)

	// result: false
	// n := int(2)

	// result: true
	// n := int(4)

	// result: true
	// n := int(6)

	// result: true
	// n := int(99856)

	// result: true
	// n := int(100000)

	// result: false
	// n := int(32897)

	// result: true
	// n := int(99999)

	// result: true
	n := int(70)

	// result: 
	// n := int(0)

	result := winnerSquareGame(n)
	fmt.Printf("result = %v\n", result)
}

