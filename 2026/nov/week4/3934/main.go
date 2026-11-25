package main

import (
	"fmt"
	"sort"
)

func smallestUniqueSubarray(nums []int) int {
	n := len(nums)
	if n == 0 {
		return 0
	}

	suffixes := make([]int, n)
	rank := make([]int, n)
	for i, value := range nums {
		suffixes[i] = i
		rank[i] = value
	}

	for length := 1; length < n; length <<= 1 {
		sort.Slice(suffixes, func(i int, j int) bool {
			a := suffixes[i]
			b := suffixes[j]
			if rank[a] != rank[b] {
				return rank[a] < rank[b]
			}

			aNext := -1
			if a+length < n {
				aNext = rank[a+length]
			}
			bNext := -1
			if b+length < n {
				bNext = rank[b+length]
			}
			return aNext < bNext
		})

		nextRank := make([]int, n)
		classes := 1
		nextRank[suffixes[0]] = 0
		for i := 1; i < n; i++ {
			previous := suffixes[i-1]
			current := suffixes[i]
			previousNext := -1
			if previous+length < n {
				previousNext = rank[previous+length]
			}
			currentNext := -1
			if current+length < n {
				currentNext = rank[current+length]
			}
			if rank[previous] != rank[current] || previousNext != currentNext {
				classes++
			}
			nextRank[current] = classes - 1
		}
		rank = nextRank
		if classes == n {
			break
		}
	}

	position := make([]int, n)
	for i, start := range suffixes {
		position[start] = i
	}
	lcp := make([]int, n)
	common := 0
	for start := range nums {
		rankIndex := position[start]
		if rankIndex == 0 {
			common = 0
			continue
		}

		other := suffixes[rankIndex-1]
		for start+common < n && other+common < n && nums[start+common] == nums[other+common] {
			common++
		}
		lcp[rankIndex] = common
		if common > 0 {
			common--
		}
	}

	result := n + 1
	for rankIndex, start := range suffixes {
		longest := 0
		if rankIndex > 0 {
			longest = max(longest, lcp[rankIndex])
		}
		if rankIndex+1 < n {
			longest = max(longest, lcp[rankIndex+1])
		}
		if longest < n-start {
			result = min(result, longest+1)
		}
	}

	if result == n+1 {
		return 0
	}
	return result
}

func main() {
	// result: 3
	// nums := []int{3,3,3}

	// result: 1
	// nums := []int{2,1,2,3,3}

	// result: 2
	nums := []int{1, 1, 2, 2, 1}

	// result: 0
	// nums := []int{}

	result := smallestUniqueSubarray(nums)
	fmt.Printf("result = %v\n", result)
}
