package main

import (
	"fmt"
	"math"
)

const maxN = int(30)

func earliestAndLatest(n int, firstPlayer int, secondPlayer int) []int {
	var earliestMemo [maxN][maxN][maxN]int
	var latestMemo [maxN][maxN][maxN]int

	if firstPlayer > secondPlayer {
		firstPlayer, secondPlayer = secondPlayer, firstPlayer
	}

	earliest, latest := dp(n, firstPlayer, secondPlayer, earliestMemo, latestMemo)
	return []int{earliest, latest}
}

func dp(n, f, s int, earliestMemo [maxN][maxN][maxN]int, latestMemo [maxN][maxN][maxN]int) (int, int) {
	if earliestMemo[n][f][s] != 0 {
		return earliestMemo[n][f][s], latestMemo[n][f][s]
	}

	if f + s == n + 1 {
		return 1, 1
	}
		
	if f + s > n + 1 {
		x, y := dp(n, n + 1 - s, n + 1 - f, earliestMemo, latestMemo)
		earliestMemo[n][f][s] = x
		latestMemo[n][f][s] = y
		return x, y
	}

	earliest := math.MaxInt32
	latest := math.MinInt32
	mid := (n + 1) / 2
		
	if s <= mid {
		for i := 0; i < f; i++ {
			for j := 0; j < s - f; j++ {
				x, y := dp(mid, i + 1, i + j + 2, earliestMemo, latestMemo)
				earliest = min(earliest, x)
				latest = max(latest, y)
			}
		}
	} else {
		sPrime := n + 1 - s
		mid2 := (n - 2 * sPrime + 1) / 2

		for i := 0; i < f; i++ {
			for j := 0; j < sPrime - f; j++ {
				x, y := dp(mid, i + 1, i + j + mid2 + 2, earliestMemo, latestMemo)
				earliest = min(earliest, x)
				latest = max(latest, y)
			}
		}
	}

	earliestMemo[n][f][s] = earliest + 1
	latestMemo[n][f][s] = latest + 1
	return earliestMemo[n][f][s], latestMemo[n][f][s]
}

func main() {
	// result: [3,4]
	// n := int(11)
	// firstPlayer := int(2)
	// secondPlayer := int(4)

	// result: [1,1]
	n := int(5)
	firstPlayer := int(1)
	secondPlayer := int(5)

	// result: []
	// n := int(0)
	// firstPlayer := int(0)
	// secondPlayer := int(0)

	result := earliestAndLatest(n, firstPlayer, secondPlayer)
	fmt.Printf("result = %v\n", result)
}

