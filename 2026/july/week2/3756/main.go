package main

import (
	"fmt"
)

const modulo = int64(1e9 + 7)

type Node struct {
	Value int64
	Sum   int64
	Count int
}

func sumAndMultiply(s string, queries [][]int) []int {
	merge := func(left, right Node, pow10 []int64) Node {
		if left.Count == 0 {
			return right
		}
		if right.Count == 0 {
			return left
		}
		return Node{
			Value: (left.Value*pow10[right.Count] + right.Value) % modulo,
			Sum:   (left.Sum + right.Sum) % modulo,
			Count: left.Count + right.Count,
		}
	}

	n := len(s)
	pow10 := make([]int64, n+1)
	pow10[0] = 1
	for i := 1; i <= n; i++ {
		pow10[i] = (pow10[i-1] * 10) % modulo
	}

	size := 1
	for size < n {
		size <<= 1
	}

	seg := make([]Node, 2*size)
	for i := 0; i < n; i++ {
		digit := int64(s[i] - '0')
		if digit == 0 {
			continue
		}
		seg[size+i] = Node{
			Value: digit,
			Sum:   digit,
			Count: 1,
		}
	}

	for i := size - 1; i >= 1; i-- {
		seg[i] = merge(seg[i<<1], seg[i<<1|1], pow10)
	}

	result := make([]int, len(queries))
	for qi, query := range queries {
		left := query[0] + size
		right := query[1] + size + 1
		leftRes := Node{}
		rightRes := Node{}

		for left < right {
			if left&1 == 1 {
				leftRes = merge(leftRes, seg[left], pow10)
				left++
			}
			if right&1 == 1 {
				right--
				rightRes = merge(seg[right], rightRes, pow10)
			}
			left >>= 1
			right >>= 1
		}

		ans := merge(leftRes, rightRes, pow10)
		result[qi] = int(ans.Value * ans.Sum % modulo)
	}

	return result
}

func main() {
	// result: [12340, 4, 9]
	// s := "10203004"
	// queries := [][]int{{0,7},{1,3},{4,6}}

	// result: [1, 0]
	// s := "1000"
	// queries := [][]int{{0,3},{1,1}}

	// result: [444444137]
	s := "9876543210"
	queries := [][]int{{0,9}}

	// result: []
	// s := ""
	// queries := [][]int{}

	result := sumAndMultiply(s, queries)
	fmt.Printf("result = %v\n", result)
}
