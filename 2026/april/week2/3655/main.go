package main

import (
	"fmt"
	"math"
)

const modulo = int64(1e9 + 7)

func xorAfterQueries(nums []int, queries [][]int) int {
	n := len(nums)
	sqrtN := int(math.Sqrt(float64(n)))
	groups := make([][][]int, sqrtN)
	for i := range sqrtN {
		groups[i] = make([][]int, 0)
	}

	for _, q := range queries {
		l, r, k, v := q[0], q[1], q[2], q[3]
		if k < sqrtN {
			groups[k] = append(groups[k], []int{l, r, v})
		} else {
			for i := l; i <= r; i += k {
				nums[i] = int((int64(nums[i]) * int64(v)) % modulo)
			}
		}
	}

	diffs := make([]int64, n + sqrtN)
	for k := 1; k < sqrtN; k++ {
		if len(groups[k]) == 0 {
			continue
		}

		for i := range diffs {
			diffs[i] = 1
		}

		for _, q := range groups[k] {
			l, r, v := q[0], q[1], q[2]
			diffs[l] = diffs[l] * int64(v) % modulo
			tmp := ((r - l) / k + 1) * k + l
			diffs[tmp] = diffs[tmp] * int64(pow(int64(v), modulo - 2)) % modulo
		}

		for i := k; i < n; i++ {
			diffs[i] = diffs[i] * diffs[i - k] % modulo
		}
		for i := range n {
			nums[i] = int((int64(nums[i]) * diffs[i]) % modulo)
		}
	}

	result := int(0)
	for _, x := range nums {
		result ^= x
	}

	return result
}

func pow(x, y int64) int {
	result := int64(1)
	for y > 0 {
		if y & 1 == 1 {
			result = result * x % modulo
		}

		x = x * x % modulo
		y >>= 1
	}

	return int(result)
}

func main() {
	// result: 4
	// nums := []int{1,1,1}
	// queries := [][]int{{0,2,1,4}}

	// result: 31
	nums := []int{2,3,1,5,4}
	queries := [][]int{{1,4,2,3},{0,2,1,2}}

	// result: 0
	// nums := []int{}
	// queries := [][]int{}

	result := xorAfterQueries(nums, queries)
	fmt.Printf("result = %v\n", result)
}

