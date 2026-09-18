package main

import (
	"fmt"
	"sort"
)

func maxNumOfSubstrings(s string) []string {
	type interval struct {
		left  int
		right int
	}
	type choice struct {
		count    int
		totalLen int
		indices  []int
	}

	first := [26]int{}
	last := [26]int{}
	for i := range first {
		first[i] = -1
		last[i] = -1
	}
	for i := range s {
		index := int(s[i] - 'a')
		if first[index] == -1 {
			first[index] = i
		}
		last[index] = i
	}

	intervals := make([]interval, 0, 26)
	for index := range first {
		if first[index] == -1 {
			continue
		}
		left, right := first[index], last[index]
		valid := true
		for i := left; i <= right; i++ {
			current := int(s[i] - 'a')
			if first[current] < left {
				valid = false
				break
			}
			if last[current] > right {
				right = last[current]
			}
		}
		if valid {
			intervals = append(intervals, interval{left: left, right: right})
		}
	}

	sort.Slice(intervals, func(i, j int) bool {
		if intervals[i].right != intervals[j].right {
			return intervals[i].right < intervals[j].right
		}
		return intervals[i].left > intervals[j].left
	})

	better := func(a, b choice) bool {
		if a.count != b.count {
			return a.count > b.count
		}
		return a.totalLen < b.totalLen
	}

	dp := make([]choice, len(intervals)+1)
	for i := 1; i <= len(intervals); i++ {
		dp[i] = dp[i-1]
		current := intervals[i-1]
		previous := sort.Search(i-1, func(j int) bool {
			return intervals[j].right >= current.left
		})
		candidate := dp[previous]
		candidate.count++
		candidate.totalLen += current.right - current.left + 1
		candidate.indices = append(append([]int{}, candidate.indices...), i-1)
		if better(candidate, dp[i]) {
			dp[i] = candidate
		}
	}

	result := make([]string, len(dp[len(intervals)].indices))
	for i, index := range dp[len(intervals)].indices {
		current := intervals[index]
		result[i] = s[current.left : current.right+1]
	}

	return result
}

func main() {
	// result: ["e","f","ccc"]
	// s := "abbaccd"

	// result: ["d","bb","cc"]
	s := "abbaccd"

	// result: []
	// s := ""

	result := maxNumOfSubstrings(s)
	fmt.Printf("result = %v\n", result)
}
