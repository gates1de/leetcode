package main

import (
	"fmt"
	"sort"
)

func findKthSmallest(coins []int, k int) int64 {
	sortedCoins := append([]int(nil), coins...)
	sort.Ints(sortedCoins)
	filtered := make([]int, 0, len(coins))

	for _, coin := range sortedCoins {
		isMultiple := false
		for _, previous := range filtered {
			if coin%previous == 0 {
				isMultiple = true
				break
			}
		}

		if !isMultiple {
			filtered = append(filtered, coin)
		}
	}

	count := func(limit int64) int64 {
		var total int64
		var search func(int, int64, int)
		search = func(start int, lcm int64, selected int) {
			for i := start; i < len(filtered); i++ {
				nextLCM := lcm / gcd(lcm, int64(filtered[i])) * int64(filtered[i])
				if nextLCM > limit {
					continue
				}

				if selected%2 == 0 {
					total += limit / nextLCM
				} else {
					total -= limit / nextLCM
				}

				search(i+1, nextLCM, selected+1)
			}
		}

		search(0, 1, 0)
		return total
	}

	left, right := int64(1), int64(filtered[0]*k)
	for left < right {
		middle := left + (right-left)/2
		if count(middle) >= int64(k) {
			right = middle
		} else {
			left = middle + 1
		}
	}

	return left
}

func gcd(a, b int64) int64 {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

func main() {
	// result: 9
	// coins := []int{3,6,9}
	// k := int(3)

	// result: 12
	coins := []int{5, 2}
	k := int(7)

	// result:
	// coins := []int{}
	// k := int(0)

	result := findKthSmallest(coins, k)
	fmt.Printf("result = %v\n", result)
}
